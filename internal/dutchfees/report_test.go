package dutchfees

import (
	"math"
	"math/big"
	"testing"
)

func TestPendingThroughTheWeek(t *testing.T) {
	cfg := networks[0]
	raw := func(n int64) *big.Int { return new(big.Int).Mul(big.NewInt(n), big.NewInt(1e18)) }
	p := protocol{Installed: cfg.Burner, RouteVerified: true, CollectorBuffer: raw(2), HookCash: raw(3), AllocatorCash: raw(4), Weight: 9000, ForwardRate: big.NewInt(1e16)}
	s := stock{collectorCash: raw(10), burnerCash: raw(5), EligibleUSD: 20, StagedRemainingUSD: 8, EligibleComplete: true, StagedRemainingComplete: true}
	pools := &poolEstimate{TargetRaw: raw(1), EligibleNonTargetUSD: 4}
	for _, tc := range []struct {
		phase    string
		mask     uint64
		expected float64
	}{
		{"sleep", 0, 33}, {"collect", 0, 33}, {"collect", 2, 25}, {"exchange", 0, 25}, {"exchange", 4, 21}, {"forward", 0, 20.86},
	} {
		p.Phase, p.GlobalMask = tc.phase, tc.mask
		got := makePending(cfg, p, s, pools, 2)
		if got.EstimateCrvUSD == nil || math.Abs(*got.EstimateCrvUSD-tc.expected) > 1e-9 || math.Abs(*got.VeEstimate-tc.expected*0.9) > 1e-9 || got.CashCrvUSD != "22" || got.Partial {
			t.Fatalf("%s mask %d: %+v", tc.phase, tc.mask, got)
		}
	}
	p.Phase = "exchange"
	p.GlobalMask = 0
	if got := makePending(cfg, p, s, nil, 0); got.EstimateCrvUSD != nil || !got.Partial {
		t.Fatal("missing target price must withhold market projection")
	}
	p.Phase = "forward"
	got := makePending(cfg, p, s, nil, 0)
	if got.EstimateCrvUSD == nil || math.Abs(*got.EstimateCrvUSD-19.87) > 1e-9 {
		t.Fatal("settled cash should be measurable without any price feed")
	}
	p.TargetMask = 8
	if got := makePending(cfg, p, s, nil, 2); got.EstimateCrvUSD != nil || got.CashCrvUSD != "22" {
		t.Fatal("FORWARD kill mask must preserve observed cash only")
	}
}

func TestPendingBridgeAndActivation(t *testing.T) {
	cfg := networks[1]
	p := protocol{Installed: cfg.Burner, Phase: "forward", RouteVerified: true, CollectorBuffer: new(big.Int), HookCash: new(big.Int), AllocatorCash: new(big.Int), ForwardRate: new(big.Int), BridgeMinimum: big.NewInt(1e18)}
	s := stock{collectorCash: big.NewInt(1e17), burnerCash: new(big.Int)}
	got := makePending(cfg, p, s, nil, 0)
	if got.VeEstimate != nil || got.BridgeMinimumMet == nil || *got.BridgeMinimumMet || *got.EstimateCrvUSD != 0.1 {
		t.Fatalf("below-minimum treasury balance: %+v", got)
	}
	p.Installed = cfg.Collector
	got = makePending(cfg, p, s, nil, 0)
	if got.EstimateCrvUSD != nil || !got.Partial || got.CashCrvUSD != "0.1" {
		t.Fatalf("pre-activation: %+v", got)
	}
}
