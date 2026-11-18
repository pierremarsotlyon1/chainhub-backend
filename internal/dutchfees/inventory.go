package dutchfees

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"main/interfaces"
)

type tokenMeta struct {
	Symbol   string
	Decimals int
	PriceUSD float64
}

func loadCatalog(ctx context.Context) ([]interfaces.CurvePool, error) {
	endpoint := os.Getenv("CURVE_POOLS_URL")
	if endpoint == "" {
		endpoint = "https://api.curve.finance/v1/getPools/all"
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{Timeout: 90 * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Curve pool API: %s", resp.Status)
	}
	var catalog interfaces.GetAllCurvePoolsResponse
	d := json.NewDecoder(resp.Body)
	d.UseNumber()
	if err := d.Decode(&catalog); err != nil {
		return nil, fmt.Errorf("Curve pool API: %w", err)
	}
	if len(catalog.Data.PoolData) == 0 {
		return nil, fmt.Errorf("Curve pool API returned empty catalog")
	}
	return catalog.Data.PoolData, nil
}

func metaNumber(v interface{}) (float64, bool) {
	f, err := strconv.ParseFloat(fmt.Sprint(v), 64)
	return f, err == nil && !math.IsInf(f, 0) && !math.IsNaN(f)
}
func catalogForChain(pools []interfaces.CurvePool, chain string) map[common.Address]tokenMeta {
	tokens := make(map[common.Address]tokenMeta)
	add := func(raw string, meta tokenMeta) {
		if !common.IsHexAddress(raw) {
			return
		}
		a := common.HexToAddress(raw)
		if a == (common.Address{}) {
			return
		}
		if prev, ok := tokens[a]; ok {
			if prev.Decimals >= 0 {
				meta.Decimals = prev.Decimals
			}
			if prev.Symbol != "" {
				meta.Symbol = prev.Symbol
			}
			if prev.PriceUSD > 0 {
				meta.PriceUSD = prev.PriceUSD
			}
		}
		tokens[a] = meta
	}
	for _, p := range pools {
		if !strings.EqualFold(p.BlockchainId, chain) {
			continue
		}
		for _, coins := range [][]interfaces.Coin{p.Coins, p.UnderlyingCoins} {
			for _, c := range coins {
				m := tokenMeta{Symbol: c.Symbol, Decimals: -1}
				if f, ok := metaNumber(c.Decimals); ok && f >= 0 && f <= 255 && math.Trunc(f) == f {
					m.Decimals = int(f)
				}
				if f, ok := metaNumber(c.UsdPrice); ok && f > 0 {
					m.PriceUSD = f
				}
				add(c.Address, m)
			}
		}
		lp := p.LpTokenAddress
		if lp == "" {
			lp = p.Address
		}
		m := tokenMeta{Decimals: 18}
		if supply, ok := metaNumber(p.TotalSupply); ok && supply > 0 && p.UsdTotal > 0 {
			price := p.UsdTotal / (supply / 1e18)
			if !math.IsInf(price, 0) {
				m.PriceUSD = price
			}
		}
		add(lp, m)
	}
	return tokens
}
func valueUSD(raw *big.Int, m tokenMeta) (float64, bool) {
	if raw == nil || m.Decimals < 0 || m.Decimals > 255 || m.PriceUSD <= 0 {
		return 0, false
	}
	f, _ := new(big.Float).SetInt(raw).Float64()
	v := f / math.Pow10(m.Decimals) * m.PriceUSD
	return v, !math.IsNaN(v) && !math.IsInf(v, 0)
}

type stockRow struct {
	StagedAvailableRaw string         `json:"stagedAvailableRaw,omitempty"`
	Token              common.Address `json:"token"`
	Symbol             string         `json:"symbol,omitempty"`
	Holder             common.Address `json:"holder"`
	Location           string         `json:"location"`
	Status             string         `json:"status"`
	Raw                string         `json:"raw"`
	Decimals           int            `json:"decimals"`
	USD                *float64       `json:"marketUSD"`
}
type stock struct {
	collectorCash, burnerCash *big.Int
	EligibleUSD               float64    `json:"eligiblePricedMarketUSD"`
	EligibleComplete          bool       `json:"eligibleValuationCompleteWithinCoverage"`
	StagedRemainingUSD        float64    `json:"stagedRemainingPricedMarketUSD"`
	StagedRemainingComplete   bool       `json:"stagedRemainingValuationComplete"`
	AssetsScanned             int        `json:"assetsScanned"`
	Unreadable                []string   `json:"unreadable,omitempty"`
	Holdings                  []stockRow `json:"holdings"`
}
type stockReader interface {
	many(context.Context, *big.Int, []call) ([]result, error)
}

func inventory(ctx context.Context, r stockReader, cfg chainConfig, block *big.Int, tokens map[common.Address]tokenMeta, installed, target common.Address, staged map[common.Address]*big.Int) (stock, error) {
	s := stock{collectorCash: new(big.Int), burnerCash: new(big.Int), EligibleComplete: true, StagedRemainingComplete: true, Holdings: []stockRow{}}
	if len(tokens) == 0 {
		return s, fmt.Errorf("%s: empty chain catalog", cfg.Name)
	}
	for a := range staged {
		if _, ok := tokens[a]; !ok {
			tokens[a] = tokenMeta{Decimals: -1}
		}
	}
	if _, ok := tokens[target]; !ok {
		tokens[target] = tokenMeta{Decimals: -1}
	}
	holders := []common.Address{cfg.Collector}
	locations := []string{"collector"}
	if installed != (common.Address{}) {
		holders = append(holders, installed)
		locations = append(locations, "installedBurner")
	}
	if cfg.Burner != installed && cfg.Burner != (common.Address{}) {
		holders = append(holders, cfg.Burner)
		locations = append(locations, "proposedBurner")
	}
	assets := make([]common.Address, 0, len(tokens))
	for a := range tokens {
		assets = append(assets, a)
	}
	sort.Slice(assets, func(i, j int) bool { return bytes.Compare(assets[i][:], assets[j][:]) < 0 })
	s.AssetsScanned = len(assets)
	type pair struct{ token, holder common.Address }
	balances := make(map[pair]*big.Int)
	var pairs []pair
	var calls []call
	for _, a := range assets {
		for _, h := range holders {
			if h == cfg.Burner && h != installed {
				if _, ok := staged[a]; !ok {
					continue
				}
			}
			c := call{a, pack("balanceOf(address)", addrWord(h))}
			if a == poolNative {
				c = call{poolMulticall, pack("getEthBalance(address)", addrWord(h))}
			}
			calls = append(calls, c)
			pairs = append(pairs, pair{a, h})
		}
	}
	read := func(calls []call) ([]*big.Int, error) {
		rows, err := r.many(ctx, block, calls)
		if err != nil {
			return nil, err
		}
		if len(rows) != len(calls) {
			return nil, fmt.Errorf("inventory: got %d of %d results", len(rows), len(calls))
		}
		out := make([]*big.Int, len(calls))
		for i, row := range rows {
			if row.OK && len(row.Data) == 32 {
				out[i] = new(big.Int).SetBytes(row.Data)
			}
		}
		return out, nil
	}
	values, err := read(calls)
	if err != nil {
		return s, err
	}
	positive := make(map[common.Address]bool)
	for i, p := range pairs {
		balances[p] = values[i]
		if values[i] != nil && values[i].Sign() > 0 {
			positive[p.token] = true
		}
	}
	calls = []call{{cfg.Collector, pack("is_killed(address)", addrWord(common.Address{}))}, {cfg.Collector, pack("epoch()")}}
	var maskAssets, decimalAssets []common.Address
	for _, a := range assets {
		if positive[a] && a != target {
			maskAssets = append(maskAssets, a)
			calls = append(calls, call{cfg.Collector, pack("is_killed(address)", addrWord(a))})
		}
	}
	for _, a := range assets {
		if positive[a] && tokens[a].Decimals < 0 {
			decimalAssets = append(decimalAssets, a)
			calls = append(calls, call{a, pack("decimals()")})
		}
	}
	values, err = read(calls)
	if err != nil {
		return s, err
	}
	global, epoch := values[0], values[1]
	if global == nil || epoch == nil {
		return s, fmt.Errorf("inventory: unreadable collector epoch or global kill mask")
	}
	for i, a := range decimalAssets {
		v := values[i+2+len(maskAssets)]
		if v != nil && v.BitLen() <= 8 {
			m := tokens[a]
			m.Decimals = int(v.Uint64())
			tokens[a] = m
		}
	}
	masks := make(map[common.Address]*big.Int)
	for i, a := range maskAssets {
		masks[a] = values[i+2]
	}
	for _, a := range assets {
		mask := masks[a]
		for j, h := range holders {
			bal, scanned := balances[pair{a, h}]
			if !scanned {
				continue
			}
			proposed := h == cfg.Burner
			initial, inWeek := staged[a]
			isStaged := proposed && installed == cfg.Burner && inWeek
			relevant := h == cfg.Collector || (proposed && h == installed)
			if bal == nil {
				s.Unreadable = append(s.Unreadable, fmt.Sprintf("%s balance of %s", a, h))
				if a == target && (h == cfg.Collector || h == installed) {
					return s, fmt.Errorf("inventory: missing target balance for %s", h)
				}
				if a != target && relevant {
					s.EligibleComplete = false
				}
				if isStaged {
					s.StagedRemainingComplete = false
				}
				continue
			}
			if bal.Sign() == 0 {
				continue
			}
			m := tokens[a]
			row := stockRow{Token: a, Symbol: m.Symbol, Holder: h, Location: locations[j], Raw: bal.String(), Decimals: m.Decimals}
			usd, priced := valueUSD(bal, m)
			if priced {
				row.USD = &usd
			}
			switch {
			case a == target:
				row.Status = "targetCash"
				if h == cfg.Collector {
					s.collectorCash.Set(bal)
				} else if h == installed {
					s.burnerCash.Set(bal)
				}
			case !relevant:
				row.Status = "otherBurnerInventory"
			case a == poolNative:
				row.Status = "nativeBalanceRequiresWrapping"
			case mask == nil:
				row.Status = "unknownKillMask"
				if relevant {
					s.EligibleComplete = false
				}
				if isStaged {
					s.StagedRemainingComplete = false
				}
			case global.Bit(2) != 0 || mask.Bit(2) != 0 || ((h == cfg.Collector || !isStaged) && (global.Bit(1) != 0 || mask.Bit(1) != 0)):
				row.Status = "excludedByKillMask"
			default:
				if isStaged {
					row.Status = "stagedOutsideExchange"
					if epoch.Uint64() == 4 {
						row.Status = "stagedInExchange"
					}
					remaining := bal
					if initial != nil && initial.Cmp(bal) < 0 {
						remaining = initial
					}
					row.StagedAvailableRaw = remaining.String()
					if value, ok := valueUSD(remaining, m); ok {
						s.StagedRemainingUSD += value
					} else {
						s.StagedRemainingComplete = false
					}
				} else if h == cfg.Collector {
					row.Status = "awaitingCollection"
				} else {
					row.Status = "unstagedBurnerBalance"
				}
				if priced {
					s.EligibleUSD += usd
				} else {
					s.EligibleComplete = false
				}
			}
			s.Holdings = append(s.Holdings, row)
		}
	}
	return s, nil
}
