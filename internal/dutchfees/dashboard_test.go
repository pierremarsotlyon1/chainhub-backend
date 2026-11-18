package dutchfees

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"math/big"
	"testing"
)

// Block 10 precedes a collection. At block 11 the same 100 crvUSD sit
// on the collector, so simulating a dispatch there must yield zero new cash.
type collectionRPC struct{}

func (collectionRPC) many(_ context.Context, block *big.Int, calls []call) ([]result, error) {
	cfg := networks[0]
	balance := pack("balanceOf(address)", addrWord(cfg.Collector))
	if len(calls) != 3 || !bytes.Equal(calls[0].Data, balance) ||
		!bytes.Equal(calls[1].Data, pack("dispatch_fees()")) || !bytes.Equal(calls[2].Data, balance) {
		return nil, fmt.Errorf("unexpected FeeSplitter probe")
	}
	before := new(big.Int)
	switch block.Uint64() {
	case 10:
	case 11:
		before.Set(wholeCrvUSD(100))
	default:
		return nil, fmt.Errorf("unrecognized snapshot block %v", block)
	}
	return []result{{true, common.LeftPadBytes(before.Bytes(), 32)}, {true, nil},
		{true, common.LeftPadBytes(wholeCrvUSD(100).Bytes(), 32)}}, nil
}

func wholeCrvUSD(n int64) *big.Int { return new(big.Int).Mul(big.NewInt(n), big.NewInt(1e18)) }

func TestCollectionBetweenDashboardSnapshots(t *testing.T) {
	cfg := networks[0]
	p := protocol{Installed: cfg.Burner, Phase: "forward", RouteVerified: true, Weight: 9000,
		CollectorBuffer: wholeCrvUSD(10), ForwardRate: big.NewInt(1e16), HookCash: new(big.Int), AllocatorCash: new(big.Int)}
	prices := map[common.Address]tokenMeta{cfg.Target: {PriceUSD: 1, Decimals: 18}}
	for _, tt := range []struct {
		name      string
		block     uint64
		collected int64
	}{{"before collection", 10, 0}, {"after collection", 11, 100}} {
		t.Run(tt.name, func(t *testing.T) {
			eth, err := readDashboardReceipts(context.Background(), collectionRPC{}, cfg, snapshot{Number: tt.block}, p, prices)
			if err != nil {
				t.Fatal(err)
			}
			estimate := makePending(cfg, p, stock{collectorCash: wholeCrvUSD(tt.collected), burnerCash: new(big.Int)}, nil, 1)
			downstream := map[string]float64{"ethereum": *estimate.VeEstimate}
			eth.beforeBuffer = estimate.beforeBuffer
			projection := Distribution{Migrated: downstream, ethereum: eth}
			published := projection.Dashboard(downstream, false)
			total := published.Components["ethereum"] + published.Components["crvusd"]
			if diff := total - 80.19; diff < -1e-9 || diff > 1e-9 || published.Partial {
				t.Fatalf("published %+v, want 80.19 USD after buffer, keeper and allocator", published)
			}
			if tt.collected > 0 && published.Components["crvusd"] != 0 {
				t.Fatal("collected fees remained in the upstream component")
			}
		})
	}
}

func TestLegacyDashboardReceiptsAndCoverage(t *testing.T) {
	d := Distribution{ethereum: &ethereumDashboard{protocol: protocol{Weight: 9000}, price: 2, controllerCrvUSD: 5, pegKeeperUSD: 20}}
	earned := map[string]float64{"ethereum": 100, "crvusd": 10, "arbitrum": 7}
	got := d.Dashboard(earned, true)
	if !got.Partial || got.Components["ethereum"] != 90 || got.Components["crvusd"] != 27 || got.Components["pegKeepers"] != 18 || got.Components["arbitrum"] != 7 || earned["crvusd"] != 10 {
		t.Fatalf("wrong published amounts or source snapshot mutated: %+v, %v", got, earned)
	}
}

func TestPausedLegacyForwardOmitsUnavailableVeCrVAmount(t *testing.T) {
	d := Distribution{ethereum: &ethereumDashboard{protocol: protocol{TargetMask: 8}, price: 1}}
	earned := map[string]float64{"ethereum": 40, "crvusd": 23, "polygon": 2}
	got := d.Dashboard(earned, false)
	if !got.Partial || len(got.Components) != 1 || got.Components["polygon"] != 2 {
		t.Fatalf("paused route: %+v", got)
	}
}
