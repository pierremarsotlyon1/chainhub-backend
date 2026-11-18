package dutchfees

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"
	"main/interfaces"
	"main/utils"
)

func installedBurner(ctx context.Context, r *reader, cfg chainConfig, head snapshot) (common.Address, error) {
	rows, err := r.many(ctx, new(big.Int).SetUint64(head.Number), []call{{cfg.Collector, pack("burner()")}})
	if err != nil {
		return common.Address{}, err
	}
	if len(rows) != 1 || !rows[0].OK || len(rows[0].Data) != 32 || !bytes.Equal(rows[0].Data[:12], make([]byte, 12)) {
		return common.Address{}, errors.New("invalid FeeCollector.burner() response")
	}
	return common.BytesToAddress(rows[0].Data), nil
}

// LegacyTargetUSD converts a legacy simulation's target-token amount at its current chain-specific price.
func LegacyTargetUSD(pools []interfaces.CurvePool, chain string, token common.Address, amount float64) (float64, bool) {
	if amount == 0 {
		return 0, true
	}
	price := catalogForChain(pools, chain)[token].PriceUSD
	usd := amount * price
	if price <= 0 || !finite(usd) {
		return 0, false
	}
	return usd, true
}

func legacyPending(cfg chainConfig, parts map[string]float64, tokens map[common.Address]tokenMeta) (*pending, error) {
	amount, ok := parts[cfg.Name]
	if !ok || !finite(amount) {
		return nil, fmt.Errorf("invalid legacy estimate for %s", cfg.Name)
	}
	price := tokens[cfg.Target].PriceUSD
	marketUSD := amount
	cash := 0.0
	if cfg.ID == 1 {
		cash, ok = parts["crvusd"]
		if !ok || !finite(cash) {
			return nil, errors.New("missing or invalid legacy crvUSD component")
		}
	} else {
		// The legacy Arbitrum estimator returns its target token balance delta (MIM).
		priceMIM := tokens[utils.TARGET_TOKEN_ARBITRUM].PriceUSD
		if amount != 0 && (priceMIM <= 0 || !finite(priceMIM)) {
			return nil, errors.New("missing MIM price for legacy Arbitrum estimate")
		}
		marketUSD *= priceMIM
	}
	value := cash
	if marketUSD != 0 {
		if price <= 0 || !finite(price) {
			return nil, errors.New("missing crvUSD price for legacy estimate")
		}
		value += marketUSD / price
	}
	if !finite(value) {
		return nil, errors.New("invalid legacy valuation")
	}
	out := &pending{EstimateCrvUSD: &value, RemainingMarketUSD: marketUSD}
	if price > 0 {
		out.TargetPriceUSD = &price
	}
	return out, nil
}

func finite(v float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) }

// MigratedDistributionFees computes only the live Dutch networks. The existing
// publisher computes the remaining networks and persists one combined result.
func MigratedDistributionFees(ctx context.Context, ethereumBlock uint64, pools []interfaces.CurvePool) (Distribution, error) {
	var d Distribution
	if ethereumBlock == 0 || len(pools) == 0 {
		return d, errors.New("distribution requires an Ethereum block and pool catalog")
	}
	rows, err := collectChains(ctx, networks, options{pools: "auto", forDistribution: true, ethereumBlock: ethereumBlock, poolCatalog: pools})
	if err != nil {
		return d, err
	}
	d.Migrated, err = distributionOverrides(rows.Chains)
	if err != nil {
		return d, err
	}
	for _, row := range rows.Chains {
		if row.Chain == "ethereum" {
			d.ethereum, d.EthereumBlockHash = row.dashboard, row.Block.Hash
		}
		if row.Pending != nil && row.Pending.Partial {
			d.partial = true
		}
	}
	if d.ethereum == nil {
		return d, errors.New("missing Ethereum fee projection")
	}
	return d, nil
}

func distributionOverrides(rows []chainReport) (map[string]float64, error) {
	values := make(map[string]float64)
	for _, row := range rows {
		if row.Status != "installed" {
			continue
		}
		if row.Pending == nil || row.Pending.EstimateCrvUSD == nil {
			return nil, fmt.Errorf("distribution estimate unavailable on %s", row.Chain)
		}
		// The existing distribution file also feeds the per-veCRV earnings/APY views.
		// L2 treasury transfers have no veCRV share on their configured routes.
		if row.Chain == "ethereum" {
			if row.Pending.VeEstimate == nil {
				return nil, errors.New("missing Ethereum allocator estimate")
			}
			value := *row.Pending.VeEstimate
			if value != 0 {
				price := row.Pending.TargetPriceUSD
				if price == nil || *price <= 0 || !finite(*price) {
					return nil, errors.New("missing crvUSD price for published Ethereum distribution")
				}
				value *= *price
			}
			values[row.Chain] = value
		} else {
			values[row.Chain] = 0
		}
	}
	return values, nil
}

func verifyBurner(ctx context.Context, r *reader, cfg chainConfig, want common.Address) error {
	head, err := r.head(ctx)
	if err != nil {
		return err
	}
	now, err := installedBurner(ctx, r, cfg, head)
	if err != nil {
		return err
	}
	if now != want {
		return fmt.Errorf("burner changed on %s during calculation; retry", cfg.Name)
	}
	return nil
}

// CheckMigration verifies the routing decision again before replacing the published file.
func CheckMigration(ctx context.Context, migrated map[string]float64) error {
	group, ctx := errgroup.WithContext(ctx)
	for _, cfg := range networks {
		group.Go(func() error {
			want := cfg.LegacyBurner
			if _, ok := migrated[cfg.Name]; ok {
				want = cfg.Burner
			}
			r, err := openReader(ctx, networkRPC(cfg), cfg.ID)
			if err != nil {
				return err
			}
			defer r.client.Close()
			return verifyBurner(ctx, r, cfg, want)
		})
	}
	return group.Wait()
}
