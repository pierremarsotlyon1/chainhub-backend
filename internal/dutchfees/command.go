package dutchfees

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
	"main/interfaces"
	"main/utils"
)

type LegacyEstimator func(context.Context, string, *ethclient.Client, []interfaces.CurvePool, bool) (map[string]float64, error)

type options struct {
	forDistribution bool
	legacy          LegacyEstimator
	chain, pools    string
	poolLimit       int
	flows           bool
	timeout         time.Duration
}

type report struct {
	WeekStart              time.Time     `json:"weekStartUTC"`
	WeekEnd                time.Time     `json:"weekEndExclusiveUTC"`
	EstimatedPendingCrvUSD *float64      `json:"estimatedPendingCrvUSD,omitempty"`
	Partial                bool          `json:"partial"`
	Chains                 []chainReport `json:"chains"`
}

// Run prints a read-only snapshot for the current Curve fee week.
func Run(ctx context.Context, args []string, stdout, stderr io.Writer, legacy LegacyEstimator) error {
	o := options{legacy: legacy}
	f := flag.NewFlagSet("weeklyPendingFees", flag.ContinueOnError)
	f.SetOutput(stderr)
	f.StringVar(&o.chain, "chain", "all", "network: all, ethereum, arbitrum, base")
	f.StringVar(&o.pools, "pools", "auto", "pool claim simulations: auto (until COLLECT closes), off or all")
	f.BoolVar(&o.flows, "flows", false, "also reconcile executed auctions and distributions over the whole week")
	f.IntVar(&o.poolLimit, "pool-limit", 0, "maximum pools simulated per network, ordered by TVL (0 = all)")
	f.DurationVar(&o.timeout, "timeout", 5*time.Minute, "deadline for the entire snapshot")
	if err := f.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	if f.NArg() != 0 || o.poolLimit < 0 || o.timeout <= 0 || (o.pools != "auto" && o.pools != "off" && o.pools != "all") {
		return errors.New("invalid arguments; see weeklyPendingFees -help")
	}
	var chains []chainConfig
	for _, cfg := range networks {
		if o.chain == "all" || cfg.Name == o.chain {
			chains = append(chains, cfg)
		}
	}
	if len(chains) == 0 {
		return fmt.Errorf("unknown network %q", o.chain)
	}

	ctx, cancel := context.WithTimeout(ctx, o.timeout)
	defer cancel()
	out, err := collectChains(ctx, chains, o)
	if err != nil {
		return err
	}
	return json.NewEncoder(stdout).Encode(out)
}

func collectChains(ctx context.Context, chains []chainConfig, o options) (report, error) {
	group, ctx := errgroup.WithContext(ctx)
	catalog := sync.OnceValues(func() ([]interfaces.CurvePool, error) { return loadCatalog(ctx) })
	out := report{Chains: make([]chainReport, len(chains))}
	for i, cfg := range chains {
		group.Go(func() error {
			row, err := readChain(ctx, cfg, o, catalog)
			if err != nil {
				return fmt.Errorf("%s: %w", cfg.Name, err)
			}
			out.Chains[i] = row
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return report{}, err
	}
	return summarizeChains(out.Chains)
}

func summarizeChains(chains []chainReport) (report, error) {
	out := report{Chains: chains}
	if len(chains) == 0 {
		return out, errors.New("no networks selected")
	}
	start := out.Chains[0].Block.Timestamp / weekSeconds * weekSeconds
	out.WeekStart = time.Unix(int64(start), 0).UTC()
	out.WeekEnd = out.WeekStart.Add(7 * 24 * time.Hour)
	var sum float64
	available := true
	covered := false
	for _, c := range out.Chains {
		if c.Block.Timestamp/weekSeconds != start/weekSeconds {
			return report{}, errors.New("networks straddle the Thursday week boundary; retry")
		}
		if c.Pending == nil {
			out.Partial = true
			continue
		}
		covered = true
		if c.Pending.EstimateCrvUSD == nil {
			available = false
		} else {
			sum += *c.Pending.EstimateCrvUSD
		}
		out.Partial = out.Partial || c.Pending.Partial
	}
	if covered && available {
		out.EstimatedPendingCrvUSD = &sum
	}
	return out, nil
}

func readChain(ctx context.Context, cfg chainConfig, o options, catalog func() ([]interfaces.CurvePool, error)) (chainReport, error) {
	out := chainReport{Chain: cfg.Name, ProposalBurner: cfg.Burner}
	r, err := openReader(ctx, networkRPC(cfg), cfg.ID)
	if err != nil {
		return out, err
	}
	defer r.client.Close()
	head, err := r.head(ctx)
	if err != nil {
		return out, err
	}
	out.Block = head
	installed, err := installedBurner(ctx, r, cfg, head)
	if err != nil {
		return out, err
	}
	out.InstalledBurner = installed
	if installed == cfg.LegacyBurner {
		out.Status = "legacy"
		if cfg.ID == 8453 {
			out.Status = "outside_legacy_coverage"
		}
		if o.legacy != nil {
			if o.flows {
				p, err := readProtocol(ctx, r, cfg, head)
				if err != nil {
					return out, err
				}
				out.Phase = p.Phase
				ledger, _, err := readWeekFlows(ctx, r, cfg, head, p)
				if err != nil {
					return out, err
				}
				out.WeekFlows = &ledger
			}
			if out.Status == "legacy" {
				pools, err := catalog()
				if err != nil {
					return out, err
				}
				components, err := o.legacy(ctx, cfg.Name, r.client, pools, false)
				if err != nil {
					return out, fmt.Errorf("legacy estimator: %w", err)
				}
				out.Pending, err = legacyPending(cfg, components, catalogForChain(pools, cfg.Name))
				if err != nil {
					return out, err
				}
				out.LegacyComponents = components
			}
		}
		if o.legacy != nil && out.LegacyComponents != nil {
			// The historical estimator reads its own latest head; catch a migration during that work.
			if err := verifyBurner(ctx, r, cfg, installed); err != nil {
				return out, err
			}
		}
		if err := r.checkHead(ctx, head); err != nil {
			return out, err
		}
		out.RPCRequests = r.requests.Load()
		return out, nil
	}
	if installed != cfg.Burner {
		return out, fmt.Errorf("unrecognized installed burner %s", installed)
	}
	p, err := readProtocol(ctx, r, cfg, head)
	if err != nil {
		return out, err
	}
	out.Phase, out.InstalledBurner, out.RouteVerified = p.Phase, p.Installed, p.RouteVerified
	out.GlobalMask, out.TargetMask = &p.GlobalMask, &p.TargetMask
	out.Status = "installed"
	if p.Installed != installed {
		return out, errors.New("burner changed at the pinned block; retry")
	}
	var staged map[common.Address]*big.Int
	if o.flows {
		var ledger flows
		ledger, staged, err = readWeekFlows(ctx, r, cfg, head, p)
		if err != nil {
			return out, err
		}
		out.WeekFlows = &ledger
	} else if p.GlobalMask&4 == 0 && (p.Phase == "collect" || p.Phase == "exchange") {
		staged, err = stagedThisWeek(ctx, r, cfg, head, p)
		if err != nil {
			if !historyUnavailable(err) {
				return out, err
			}
			pools, catErr := catalog()
			if catErr != nil {
				return out, catErr
			}
			tokens := catalogForChain(pools, cfg.Name)
			staged, err = catalogStages(ctx, r, cfg, new(big.Int).SetUint64(head.Number), tokens, p.CollectStart, p.CollectEnd)
			if err != nil {
				return out, err
			}
			out.StageSource = "catalogContractReads"
		} else {
			out.StageSource = "events"
		}
	}
	block := new(big.Int).SetUint64(head.Number)
	early := beforeCollectCloses(p)
	scanPools := o.pools == "all" || (o.pools == "auto" && early)
	fullInventory := early && p.GlobalMask&6 == 0
	stagedInventory := (p.Phase == "collect" || p.Phase == "exchange") && len(staged) > 0 && p.GlobalMask&4 == 0
	tokens := map[common.Address]tokenMeta{p.Target: {Symbol: "crvUSD", Decimals: 18}}
	out.InventoryScope = "crvUSD"
	var prices map[common.Address]tokenMeta
	var pools []interfaces.CurvePool
	if scanPools || fullInventory || stagedInventory || (o.forDistribution && cfg.ID == 1) {
		pools, err = catalog()
		if err != nil {
			return out, err
		}
		prices = catalogForChain(pools, cfg.Name)
		if len(prices) == 0 {
			return out, errors.New("empty Curve catalog for network")
		}
		if meta, ok := prices[p.Target]; ok {
			tokens[p.Target] = meta
		}
	}
	if fullInventory {
		tokens = prices
		out.InventoryScope = "curveCatalogAndWeeklyLots"
	} else if stagedInventory {
		out.InventoryScope = "crvUSDAndWeeklyLots"
		for a := range staged {
			if meta, ok := prices[a]; ok {
				tokens[a] = meta
			}
		}
	}
	var inventoryLots map[common.Address]*big.Int
	if fullInventory || stagedInventory {
		inventoryLots = staged
	}
	stock, err := inventory(ctx, r, cfg, block, tokens, p.Installed, p.Target, inventoryLots)
	if err != nil {
		return out, err
	}
	out.Inventory = &stock
	if scanPools {
		estimate, err := poolAccrual(ctx, r, cfg, block, pools, prices, o.poolLimit, p.Target)
		if err != nil {
			return out, err
		}
		out.Pools = &estimate
	}
	estimate := makePending(cfg, p, stock, out.Pools, tokens[p.Target].PriceUSD)
	out.Pending = &estimate
	if out.StageSource == "catalogContractReads" {
		out.Pending.Partial = true
	}
	if stagedInventory && p.Phase == "exchange" {
		out.Pending.CurrentLotQuote, err = auctionQuote(ctx, r, block, cfg.Burner, staged)
		if err != nil {
			return out, err
		}
	}
	if err = r.checkHead(ctx, head); err != nil {
		return out, err
	}
	out.RPCRequests = r.requests.Load()
	return out, nil
}

// Staging is restricted to the collector's half-open COLLECT interval.
func stagedThisWeek(ctx context.Context, r *reader, cfg chainConfig, head snapshot, p protocol) (map[common.Address]*big.Int, error) {
	from, err := r.blockAtStart(ctx, head, p.CollectStart, cfg)
	if err != nil {
		return nil, err
	}
	to := head.Number
	if p.Phase == "exchange" {
		after, err := r.blockAtStart(ctx, head, p.CollectEnd, cfg)
		if err != nil {
			return nil, err
		}
		if after <= from {
			return nil, nil
		}
		to = after - 1
	}
	return weeklyStages(ctx, r, cfg, from, to)
}

func readWeekFlows(ctx context.Context, r *reader, cfg chainConfig, head snapshot, p protocol) (flows, map[common.Address]*big.Int, error) {
	from, err := r.blockAtStart(ctx, head, head.Timestamp/weekSeconds*weekSeconds, cfg)
	if err != nil {
		return flows{}, nil, err
	}
	ledger, staged, err := weeklyFlows(ctx, r, cfg, from, head.Number, p.Target, p.Hooker)
	if err != nil {
		return flows{}, nil, fmt.Errorf("weekly event history (check %s_RPC_URL): %w", strings.ToUpper(cfg.Name), err)
	}
	return ledger, staged, nil
}

func networkRPC(cfg chainConfig) string {
	endpoint := os.Getenv(strings.ToUpper(cfg.Name) + "_RPC_URL")
	if endpoint == "" {
		name := cfg.Name
		if cfg.ID == 1 {
			name = "mainnet"
		}
		endpoint = utils.GetPublicRpcUrl(name)
	}
	return endpoint
}
