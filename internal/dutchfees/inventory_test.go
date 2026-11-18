package dutchfees

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"main/interfaces"
	"math/big"
	"testing"
)

func TestCatalogPerChainAndLPPrice(t *testing.T) {
	a, b, lp := common.HexToAddress("0x01"), common.HexToAddress("0x02"), common.HexToAddress("0x03")
	pools := []interfaces.CurvePool{{BlockchainId: "base", LpTokenAddress: lp.Hex(), UsdTotal: 12, TotalSupply: "2000000000000000000", Coins: []interfaces.Coin{{Address: a.Hex(), Decimals: "6", UsdPrice: nil}}, UnderlyingCoins: []interfaces.Coin{{Address: a.Hex(), UsdPrice: "1.01"}}}, {BlockchainId: "ethereum", Coins: []interfaces.Coin{{Address: b.Hex()}}}}
	tokens := catalogForChain(pools, "base")
	if len(tokens) != 2 || tokens[a].Decimals != 6 || tokens[a].PriceUSD != 1.01 || tokens[lp].PriceUSD != 6 {
		t.Fatal(tokens)
	}
}

type stockStub map[string]int64

func (stub stockStub) many(_ context.Context, _ *big.Int, calls []call) ([]result, error) {
	out := make([]result, len(calls))
	for i, c := range calls {
		n := stub[c.To.Hex()+string(c.Data)]
		out[i] = result{OK: true, Data: common.LeftPadBytes(big.NewInt(n).Bytes(), 32)}
	}
	return out, nil
}
func TestInventoryCapsStagedAndExcludesCashAndKilled(t *testing.T) {
	cfg := chainConfig{Collector: common.HexToAddress("0xCA"), Burner: common.HexToAddress("0xBA")}
	lot, target, killed := common.HexToAddress("0x01"), common.HexToAddress("0x02"), common.HexToAddress("0x03")
	stub := stockStub{}
	set := func(a common.Address, data []byte, n int64) { stub[a.Hex()+string(data)] = n }
	set(cfg.Collector, pack("epoch()"), 4)
	set(cfg.Collector, pack("is_killed(address)", addrWord(killed)), 2)
	set(target, pack("balanceOf(address)", addrWord(cfg.Collector)), 53)
	set(lot, pack("balanceOf(address)", addrWord(cfg.Burner)), 20)
	set(killed, pack("balanceOf(address)", addrWord(cfg.Collector)), 80)
	meta := map[common.Address]tokenMeta{lot: {Decimals: 0, PriceUSD: 1}, target: {Decimals: 0}, killed: {Decimals: 0, PriceUSD: 1}}
	got, err := inventory(context.Background(), stub, cfg, big.NewInt(1), meta, cfg.Burner, target, map[common.Address]*big.Int{lot: big.NewInt(10)})
	if err != nil || !got.EligibleComplete || !got.StagedRemainingComplete || got.EligibleUSD != 20 || got.StagedRemainingUSD != 10 || got.collectorCash.Int64() != 53 {
		t.Fatal(got, err)
	}
}

func TestInventoryBeforeBurnerReplacement(t *testing.T) {
	cfg := chainConfig{Collector: common.HexToAddress("0xCA"), Burner: common.HexToAddress("0xBA")}
	old, lot, target := common.HexToAddress("0xBB"), common.HexToAddress("0x01"), common.HexToAddress("0x02")
	stub := stockStub{}
	for _, h := range []common.Address{old, cfg.Burner, cfg.Collector} {
		stub[lot.Hex()+string(pack("balanceOf(address)", addrWord(h)))] = 20
	}
	meta := map[common.Address]tokenMeta{lot: {Decimals: 0, PriceUSD: 1}}
	got, err := inventory(context.Background(), stub, cfg, big.NewInt(1), meta, old, target, map[common.Address]*big.Int{lot: big.NewInt(10)})
	if err != nil || len(got.Holdings) != 3 || got.EligibleUSD != 20 || got.StagedRemainingUSD != 0 {
		t.Fatal(got, err)
	}
}
