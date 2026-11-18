package dutchfees

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	mc "main/contracts/multicall"
	"main/interfaces"
	"main/utils"
)

var poolMulticall = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
var poolNative = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

type poolEstimate struct {
	EligibleUSD          float64  `json:"eligiblePricedClaimableUSD"`
	EligibleNonTargetUSD float64  `json:"eligiblePricedNonTargetUSD"`
	TargetCrvUSD         string   `json:"claimableCrvUSD"`
	TargetRaw            *big.Int `json:"-"`
	KilledAssets         int      `json:"individuallyKilledAssets"`
	UnknownMasks         int      `json:"unknownEligibilityAssets"`
	USD                  float64  `json:"pricedClaimableUSD"`
	Partial              bool     `json:"partial"`
	Listed               int      `json:"uniqueListedPools"`
	Invalid              int      `json:"invalidCatalogueRows"`
	Attempted            int      `json:"attemptedPools"`
	Omitted              int      `json:"omittedByLimit"`
	WithReceipt          int      `json:"poolsWithCollectorReceipt"`
	NoReceipt            int      `json:"successfulClaimWithoutVerifiedReceipt"`
	Failed               int      `json:"unsupportedOrFailedPools"`
	Unreadable           int      `json:"unreadableTokenBalancePairs"`
	Unpriced             int      `json:"unpricedPositiveReceipts"`
}

func listedPools(pools []interfaces.CurvePool, chain string, limit int) ([]interfaces.CurvePool, poolEstimate) {
	var out poolEstimate
	var selected []interfaces.CurvePool
	seen := map[common.Address]bool{}
	for _, p := range pools {
		if !strings.EqualFold(p.BlockchainId, chain) {
			continue
		}
		addr := common.HexToAddress(p.Address)
		if !common.IsHexAddress(p.Address) || addr == (common.Address{}) {
			out.Invalid++
			continue
		}
		if seen[addr] {
			continue
		}
		seen[addr] = true
		selected = append(selected, p)
	}
	sort.SliceStable(selected, func(i, j int) bool { return selected[i].UsdTotal > selected[j].UsdTotal })
	out.Listed = len(selected)
	if limit > 0 && len(selected) > limit {
		out.Omitted = len(selected) - limit
		selected = selected[:limit]
	}
	out.Partial = out.Invalid > 0 || out.Listed == 0 || out.Omitted > 0
	return selected, out
}

func poolAssets(p interfaces.CurvePool) ([]common.Address, int) {
	var assets []common.Address
	invalid := 0
	seen := map[common.Address]bool{}
	add := func(raw string) {
		a := common.HexToAddress(raw)
		if !common.IsHexAddress(raw) || a == (common.Address{}) {
			invalid++
			return
		}
		if !seen[a] {
			seen[a] = true
			assets = append(assets, a)
		}
	}
	for _, c := range p.Coins {
		add(c.Address)
	}
	if p.LpTokenAddress != "" {
		add(p.LpTokenAddress)
	} else {
		add(p.Address)
	}
	return assets, invalid
}

// tryAggregate executes its calls in order inside a single eth_call. This observes
// the collector's exact receipt, including LP tokens minted by crypto pools.
func poolSimulation(a *abi.ABI, p interfaces.CurvePool, collector common.Address, assets []common.Address) (call, error) {
	balances := make([]mc.Multicall3Call, 0, len(assets))
	for _, asset := range assets {
		q := mc.Multicall3Call{Target: asset, CallData: pack("balanceOf(address)", addrWord(collector))}
		if asset == poolNative {
			q.Target = poolMulticall
			q.CallData = pack("getEthBalance(address)", addrWord(collector))
		}
		balances = append(balances, q)
	}
	pool := common.HexToAddress(p.Address)
	inner := append([]mc.Multicall3Call{}, balances...)
	// FeeCollector.withdraw_many forwards the pool call as the collector.
	inner = append(inner, mc.Multicall3Call{Target: collector, CallData: pack("withdraw_many(address[])", uintWord(32), uintWord(1), addrWord(pool))})
	if strings.EqualFold(p.BlockchainId, "ethereum") {
		// The legacy proxy can withdraw as pool admin. Its separate burn step
		// requires an EOA. Only a receipt already credited to the collector counts.
		inner = append(inner, mc.Multicall3Call{Target: utils.POOL_OWNER_PENDING_FEES_MAINNET, CallData: pack("withdraw_admin_fees(address)", addrWord(pool))})
	}
	for _, sig := range []string{"withdraw_admin_fees()", "claim_admin_fees()"} {
		inner = append(inner, mc.Multicall3Call{Target: pool, CallData: pack(sig)})
	}
	inner = append(inner, balances...)
	data, err := a.Pack("tryAggregate", false, inner)
	return call{To: poolMulticall, Data: data}, err
}

func recordPool(out *poolEstimate, rows []mc.Multicall3Result, assets []common.Address, tokens map[common.Address]tokenMeta, receipts map[common.Address]*big.Int) {
	n := len(assets)
	claimCount := len(rows) - 2*n
	if claimCount != 3 && claimCount != 4 {
		out.Failed++
		out.Partial = true
		return
	}
	succeeded := false
	for _, claim := range rows[n : n+claimCount] {
		succeeded = succeeded || claim.Success
	}
	if !succeeded {
		out.Failed++
		out.Partial = true
		return
	}
	received := false
	for i, asset := range assets {
		before, after := rows[i], rows[n+claimCount+i]
		if !before.Success || !after.Success || len(before.ReturnData) != 32 || len(after.ReturnData) != 32 {
			out.Unreadable++
			out.Partial = true
			continue
		}
		delta := new(big.Int).Sub(new(big.Int).SetBytes(after.ReturnData), new(big.Int).SetBytes(before.ReturnData))
		if delta.Sign() < 0 {
			out.Unreadable++
			out.Partial = true
			continue
		}
		if delta.Sign() == 0 {
			continue
		}
		received = true
		if receipts[asset] == nil {
			receipts[asset] = new(big.Int)
		}
		receipts[asset].Add(receipts[asset], delta)
		meta, ok := tokens[asset]
		usd, priced := valueUSD(delta, meta)
		if !ok || !priced {
			out.Unpriced++
			out.Partial = true
			continue
		}
		out.USD += usd
	}
	if received {
		out.WithReceipt++
	} else {
		out.NoReceipt++
		out.Partial = true
	}
}

func poolAccrual(ctx context.Context, r *reader, cfg chainConfig, block *big.Int, pools []interfaces.CurvePool, tokens map[common.Address]tokenMeta, limit int, target common.Address) (poolEstimate, error) {
	selected, out := listedPools(pools, cfg.Name, limit)
	receipts := make(map[common.Address]*big.Int)
	out.TargetRaw = new(big.Int)
	out.TargetCrvUSD = "0"
	a, err := mc.MulticallMetaData.GetAbi()
	if err != nil {
		return out, err
	}
	const chunk = 8
	for start := 0; start < len(selected); start += chunk {
		end := start + chunk
		if end > len(selected) {
			end = len(selected)
		}
		var calls []call
		var byPool [][]common.Address
		for _, p := range selected[start:end] {
			assets, invalid := poolAssets(p)
			out.Unreadable += invalid
			if invalid > 0 {
				out.Partial = true
			}
			q, err := poolSimulation(a, p, cfg.Collector, assets)
			if err != nil {
				out.Partial = true
				return out, err
			}
			calls = append(calls, q)
			byPool = append(byPool, assets)
		}
		out.Attempted += len(calls)
		results, err := r.many(ctx, block, calls)
		if err != nil {
			out.Partial = true
			return out, fmt.Errorf("pool simulation batch at pool %d: %w", start, err)
		}
		if len(results) != len(calls) {
			out.Partial = true
			return out, fmt.Errorf("pool simulation received %d results for %d pools", len(results), len(calls))
		}
		for i, res := range results {
			if !res.OK {
				out.Failed++
				out.Partial = true
				continue
			}
			decoded, err := a.Unpack("tryAggregate", res.Data)
			if err != nil || len(decoded) != 1 {
				out.Failed++
				out.Partial = true
				continue
			}
			rows := *abi.ConvertType(decoded[0], new([]mc.Multicall3Result)).(*[]mc.Multicall3Result)
			recordPool(&out, rows, byPool[i], tokens, receipts)
		}
	}
	err = poolEligibility(ctx, r, cfg.Collector, block, target, receipts, tokens, &out)
	return out, err
}

func poolEligibility(ctx context.Context, r stockReader, collector common.Address, block *big.Int, target common.Address, receipts map[common.Address]*big.Int, tokens map[common.Address]tokenMeta, out *poolEstimate) error {
	out.TargetRaw = new(big.Int)
	if amount := receipts[target]; amount != nil {
		out.TargetRaw.Set(amount)
		if usd, ok := valueUSD(amount, tokens[target]); ok {
			out.EligibleUSD += usd
		}
	}
	out.TargetCrvUSD = amountText(out.TargetRaw, 18)
	var assets []common.Address
	for a := range receipts {
		if a != target && a != poolNative {
			assets = append(assets, a)
		}
	}
	if len(assets) == 0 {
		return nil
	}
	sort.Slice(assets, func(i, j int) bool { return assets[i].Hex() < assets[j].Hex() })
	calls := make([]call, len(assets))
	for i, a := range assets {
		calls[i] = call{collector, pack("is_killed(address)", addrWord(a))}
	}
	masks, err := r.many(ctx, block, calls)
	if err != nil || len(masks) != len(calls) {
		out.Partial = true
		out.UnknownMasks = len(calls)
		if err != nil {
			return fmt.Errorf("pool asset eligibility: %w", err)
		}
		return fmt.Errorf("pool asset eligibility: got %d of %d results", len(masks), len(calls))
	}
	for i, a := range assets {
		m := masks[i]
		if !m.OK || len(m.Data) != 32 {
			out.UnknownMasks++
			out.Partial = true
			continue
		}
		mask := new(big.Int).SetBytes(m.Data)
		if mask.Bit(1) != 0 || mask.Bit(2) != 0 {
			out.KilledAssets++
			continue
		}
		if usd, ok := valueUSD(receipts[a], tokens[a]); ok {
			out.EligibleNonTargetUSD += usd
			out.EligibleUSD += usd
		}
	}
	return nil
}
