package src

import (
	"testing"

	"main/interfaces"
)

func TestLegacyDashboardTargetConversion(t *testing.T) {
	pools := []interfaces.CurvePool{{BlockchainId: ARBITRUM_CHAIN, Coins: []interfaces.Coin{
		{Address: TARGET_TOKEN[ARBITRUM_CHAIN].Hex(), UsdPrice: 0.75},
	}}}
	earned := map[string]float64{"ethereum": 99, "crvusd": 5, "arbitrum": 8, "polygon": 40, "fantom": 0, "optimism": 3}
	converted, partial := legacyDashboardUSD(pools, earned, map[string]float64{"optimism": 3})
	if !partial || converted["arbitrum"] != 6 || earned["arbitrum"] != 8 || converted["ethereum"] != 99 || converted["crvusd"] != 5 || converted["optimism"] != 3 {
		t.Fatalf("converted %v, source %v, partial %t", converted, earned, partial)
	}
	if _, ok := converted["polygon"]; ok {
		t.Fatal("unpriced target token was reported in USD")
	}
	if value, ok := converted["fantom"]; !ok || value != 0 {
		t.Fatal("zero target amount required a price")
	}
}
