package dutchfees

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"main/utils"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	mc "main/contracts/multicall"
	"main/interfaces"
)

func TestPoolSimulationUsesOnlyIncrementalCollectorReceipt(t *testing.T) {
	token := common.HexToAddress("0x1234")
	balance := func(n int64) mc.Multicall3Result {
		return mc.Multicall3Result{Success: true, ReturnData: common.LeftPadBytes(big.NewInt(n).Bytes(), 32)}
	}
	rows := []mc.Multicall3Result{balance(500), {Success: true}, {}, {}, balance(503)}
	meta := map[common.Address]tokenMeta{token: {PriceUSD: 2, Decimals: 0}}
	receipts := map[common.Address]*big.Int{}
	var out poolEstimate
	recordPool(&out, rows, []common.Address{token}, meta, receipts)
	if out.Partial || out.USD != 6 || out.WithReceipt != 1 {
		t.Fatal(out)
	}
	rows[4] = balance(500)
	recordPool(&out, rows, []common.Address{token}, meta, receipts)
	if !out.Partial || out.NoReceipt != 1 || out.USD != 6 {
		t.Fatal(out)
	}
}

func TestListedPoolsBudget(t *testing.T) {
	a, b := common.HexToAddress("0x12").Hex(), common.HexToAddress("0x13").Hex()
	pools := []interfaces.CurvePool{{BlockchainId: "base", Address: a, UsdTotal: 1}, {BlockchainId: "base", Address: b, UsdTotal: 2}, {BlockchainId: "base", Address: a}, {BlockchainId: "ethereum", Address: a}}
	selected, out := listedPools(pools, "base", 1)
	if len(selected) != 1 || selected[0].Address != b || out.Omitted != 1 || !out.Partial {
		t.Fatal(selected, out)
	}
	selected, out = listedPools(pools, "base", 0)
	if len(selected) != 2 || out.Omitted != 0 || out.Partial {
		t.Fatal(selected, out)
	}
}

func TestPoolEligibilityKeepsTargetRawAndExcludesKilled(t *testing.T) {
	collector := common.HexToAddress("0xCA")
	lot, target, killed := common.HexToAddress("0x01"), common.HexToAddress("0x02"), common.HexToAddress("0x03")
	stub := stockStub{collector.Hex() + string(pack("is_killed(address)", addrWord(killed))): 4}
	raw := map[common.Address]*big.Int{lot: big.NewInt(2), target: big.NewInt(9), killed: big.NewInt(100), poolNative: big.NewInt(1000)}
	meta := map[common.Address]tokenMeta{lot: {Decimals: 0, PriceUSD: 3}, target: {Decimals: 18}, killed: {Decimals: 0, PriceUSD: 1}, poolNative: {Decimals: 0, PriceUSD: 9999}}
	var got poolEstimate
	err := poolEligibility(context.Background(), stub, collector, big.NewInt(1), target, raw, meta, &got)
	if err != nil || got.EligibleUSD != 6 || got.EligibleNonTargetUSD != 6 || got.TargetRaw.Cmp(big.NewInt(9)) != 0 || got.KilledAssets != 1 {
		t.Fatal(got, err)
	}
}

func TestPoolSimulationRoutes(t *testing.T) {
	contractABI, err := mc.MulticallMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}
	pool, collector := common.HexToAddress("0x1234"), common.HexToAddress("0x4567")
	for _, chain := range []string{"ethereum", "arbitrum", "base"} {
		q, err := poolSimulation(contractABI, interfaces.CurvePool{Address: pool.Hex(), BlockchainId: chain}, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		args, err := contractABI.Methods["tryAggregate"].Inputs.Unpack(q.Data[4:])
		if err != nil {
			t.Fatal(err)
		}
		calls := *abi.ConvertType(args[1], new([]mc.Multicall3Call)).(*[]mc.Multicall3Call)
		n := 3
		if chain == "ethereum" {
			n = 4
		}
		if len(calls) != n || calls[0].Target != collector || !bytes.Equal(calls[0].CallData, pack("withdraw_many(address[])", uintWord(32), uintWord(1), addrWord(pool))) {
			t.Fatal(chain, calls)
		}
		if chain == "ethereum" && (calls[1].Target != utils.POOL_OWNER_PENDING_FEES_MAINNET || !bytes.Equal(calls[1].CallData, pack("withdraw_admin_fees(address)", addrWord(pool)))) {
			t.Fatal(calls[1])
		}
		if calls[n-1].Target != pool || calls[n-2].Target != pool {
			t.Fatal(calls)
		}
	}
}
