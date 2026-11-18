package dutchfees

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"main/interfaces"
	"main/utils"
)

func serveMigrationBurner(t *testing.T, cfg chainConfig, burner common.Address) {
	t.Helper()
	server := rpc.NewServer()
	t.Cleanup(server.Stop)
	s := &commandRPC{cfg: cfg, head: 7, replies: map[string][]byte{
		cfg.Collector.Hex() + string(pack("burner()")): addrWord(burner),
	}}
	if err := server.RegisterName("eth", s); err != nil {
		t.Fatal(err)
	}
	h := httptest.NewServer(server)
	t.Cleanup(h.Close)
	t.Setenv(strings.ToUpper(cfg.Name)+"_RPC_URL", h.URL)
}

func TestMigrationRoutesAndSelectedTotals(t *testing.T) {
	var rows []chainReport
	for _, cfg := range networks {
		serveMigrationBurner(t, cfg, cfg.LegacyBurner)
		calls := 0
		opts := options{legacy: func(ctx context.Context, chain string, client *ethclient.Client, pools []interfaces.CurvePool, updateCaches bool) (map[string]float64, error) {
			calls++
			if updateCaches || chain != cfg.Name || client == nil || len(pools) != 1 || pools[0].BlockchainId != chain {
				t.Fatalf("unexpected legacy estimator inputs on %s with updateCaches=%v", chain, updateCaches)
			}
			if chain == "ethereum" {
				return map[string]float64{chain: 12, "crvusd": 5}, nil
			}
			return map[string]float64{chain: 8}, nil
		}}
		row, err := readChain(context.Background(), cfg, opts, func() ([]interfaces.CurvePool, error) {
			if cfg.Name == "base" {
				t.Fatal("uncovered Base fetched the legacy catalog")
			}
			return []interfaces.CurvePool{{BlockchainId: cfg.Name, Coins: []interfaces.Coin{
				{Address: cfg.Target.Hex(), Decimals: 18, UsdPrice: 2},
				{Address: utils.TARGET_TOKEN_ARBITRUM.Hex(), Decimals: 18, UsdPrice: 0.5},
			}}}, nil
		})
		if err != nil {
			t.Fatal(err)
		}
		rows = append(rows, row)
		if cfg.Name == "base" {
			if row.Status != "outside_legacy_coverage" || calls != 0 || row.Pending != nil {
				t.Fatalf("incorrect Base fallback %+v", row)
			}
			continue
		}
		expected := map[string]float64{"ethereum": 11, "arbitrum": 2}[cfg.Name]
		if row.Status != "legacy" || calls != 1 || row.Pending == nil || row.Pending.EstimateCrvUSD == nil || *row.Pending.EstimateCrvUSD != expected {
			t.Fatalf("incorrect %s estimate %+v", cfg.Name, row.Pending)
		}
	}

	for _, tc := range []struct {
		selection        []chainReport
		want             float64
		partial, missing bool
	}{
		{rows[:1], 11, false, false}, {rows[:2], 13, false, false}, {rows, 13, true, false}, {rows[2:], 0, true, true},
	} {
		got, err := summarizeChains(tc.selection)
		if err != nil {
			t.Fatal(err)
		}
		if got.Partial != tc.partial || (got.EstimatedPendingCrvUSD == nil) != tc.missing || (!tc.missing && *got.EstimatedPendingCrvUSD != tc.want) {
			t.Fatalf("unexpected summary %+v", got)
		}
	}
}

func TestMigrationRejectsUnknownBurner(t *testing.T) {
	cfg := networks[0]
	serveMigrationBurner(t, cfg, address("0x1234"))
	_, err := readChain(context.Background(), cfg, options{}, nil)
	if err == nil || !strings.Contains(err.Error(), "unrecognized installed burner") {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestDistributionOverrides(t *testing.T) {
	gross, ve, price := 100.0, 65.0, 0.98
	rows := []chainReport{
		{Chain: "ethereum", Status: "installed", Pending: &pending{EstimateCrvUSD: &gross, VeEstimate: &ve, TargetPriceUSD: &price}},
		{Chain: "arbitrum", Status: "legacy"},
		{Chain: "base", Status: "installed", Pending: &pending{EstimateCrvUSD: &gross}},
	}
	got, err := distributionOverrides(rows)
	if err != nil {
		t.Fatal(err)
	}
	base, ok := got["base"]
	if len(got) != 2 || got["ethereum"] != ve*price || !ok || base != 0 {
		t.Fatalf("unexpected allocation %v", got)
	}
	rows[0].Pending.VeEstimate = nil
	if _, err := distributionOverrides(rows); err == nil {
		t.Fatal("missing veCRV share accepted")
	}
}
