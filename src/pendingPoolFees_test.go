package src

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"main/interfaces"
	"main/utils"
)

// Run explicitly from a backend checkout with network access and the historic simulation credentials.
func TestLiveLegacyPendingFees(t *testing.T) {
	if os.Getenv("CHECK_LIVE_LEGACY_FEES") != "1" {
		t.Skip("live network check")
	}
	if workdir := os.Getenv("LEGACY_BACKEND_WORKDIR"); workdir != "" {
		before, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			if err := os.Chdir(before); err != nil {
				t.Error(err)
			}
		})
		if err := os.Chdir(workdir); err != nil {
			t.Fatal(err)
		}
	}
	pools, err := utils.GetAllCurvePools()
	if err != nil {
		t.Fatal(err)
	}
	for _, chain := range []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN} {
		t.Run(chain, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel()
			var client *ethclient.Client
			if rpcURL := os.Getenv("LEGACY_TEST_RPC_URL"); rpcURL != "" {
				var err error
				client, err = ethclient.DialContext(ctx, rpcURL)
				if err != nil {
					t.Fatal("cannot connect test RPC")
				}
				defer client.Close()
			}
			fees, err := LegacyPendingFees(ctx, chain, client, pools, false)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("fees %v", fees)
		})
	}
}

func TestLegacyCurveCatalogNumber(t *testing.T) {
	if price := legacyUSDPrice(json.Number("0.975")); price != 0.975 {
		t.Fatalf("price %v", price)
	}
	for _, input := range []any{"18", 18, 18.0, json.Number("18")} {
		if got := legacyDecimals(input); got != 18 {
			t.Errorf("decimals %v: %d", input, got)
		}
	}
	if legacyDecimals(json.Number("1.5")) != -1 {
		t.Fatal("accepted fractional decimals")
	}
}

func TestSimulatedBalanceRequiresDecodedUnsignedInteger(t *testing.T) {
	if _, err := simulatedBalance(interfaces.Result{}); err == nil {
		t.Fatal("accepted missing trace")
	}
	for _, tc := range []struct {
		value any
		valid bool
	}{{"0", true}, {"32768000000000000000", true}, {"", false}, {"-1", false}, {nil, false}} {
		trace := interfaces.Trace{DecodedOutput: []interfaces.DecodedOutput{{Value: tc.value}}}
		balance, err := simulatedBalance(interfaces.Result{Trace: []interfaces.Trace{trace}})
		if (err == nil) != tc.valid {
			t.Errorf("value %v: err %v", tc.value, err)
		}
		if tc.valid && balance.String() != tc.value {
			t.Errorf("value %v: balance %s", tc.value, balance)
		}
	}
}
