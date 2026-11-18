package src

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
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

func TestLegacyCutoffBlockCapsHistoricalReads(t *testing.T) {
	for _, tc := range []struct {
		cutoff, lookupHeight, want uint64
		lookupCalled               bool
	}{{80, 7, 7, true}, {80, 15, 10, true}, {80, 0, 0, true}, {100, 15, 10, false}, {101, 15, 10, false}} {
		called := false
		got := legacyCutoffBlock("ethereum", tc.cutoff, 100, 10, func(_ string, timestamp uint64) uint64 {
			called = true
			if timestamp != tc.cutoff {
				t.Fatalf("unexpected cutoff %d", timestamp)
			}
			return tc.lookupHeight
		})
		if got != tc.want || called != tc.lookupCalled {
			t.Errorf("cutoff %d, lookup height %d: got %d, lookup called %v", tc.cutoff, tc.lookupHeight, got, called)
		}
	}
}

type legacyPinnedRPC struct {
	tags []string
}

func (s *legacyPinnedRPC) GetBlockByNumber(tag string, _ bool) (*types.Header, error) {
	s.tags = append(s.tags, tag)
	return &types.Header{Number: big.NewInt(42), Difficulty: big.NewInt(1), Time: 200}, nil
}

func (s *legacyPinnedRPC) Call(_ map[string]any, tag string) (hexutil.Bytes, error) {
	s.tags = append(s.tags, tag)
	switch len(s.tags) {
	case 2:
		return make([]byte, 32), nil
	case 3:
		return make([]byte, 64), nil
	default:
		return nil, errors.New("stop after forward epoch read")
	}
}

func TestLegacyPendingFeesAtPinsCollectorReads(t *testing.T) {
	server := rpc.NewServer()
	defer server.Stop()
	stub := &legacyPinnedRPC{}
	if err := server.RegisterName("eth", stub); err != nil {
		t.Fatal(err)
	}
	client := ethclient.NewClient(rpc.DialInProc(server))
	defer client.Close()
	pools := []interfaces.CurvePool{{BlockchainId: ETHEREUM_CHAIN}}
	_, err := LegacyPendingFeesAt(context.Background(), ETHEREUM_CHAIN, client, pools, false, big.NewInt(42))
	if err == nil || !strings.Contains(err.Error(), "stop after forward epoch read") {
		t.Fatalf("expected forward epoch call, got %v", err)
	}
	if !slices.Equal(stub.tags, []string{"0x2a", "0x2a", "0x2a", "0x2a"}) {
		t.Fatalf("header, burner and epoch block tags %v", stub.tags)
	}
}

func TestLegacyPostCollectTransferDoesNotReduceEarlierCash(t *testing.T) {
	var receipts legacySplitterReceipts
	receipts.add(99, 100, big.NewInt(60))
	receipts.add(100, 100, big.NewInt(40))
	receipts.add(101, 100, big.NewInt(1000))
	// The collector held 250 at the cutoff: 100 from the splitter, 150 from
	// elsewhere. The later 1000 only changes the week's splitter component.
	if other := receipts.otherCashAtCollect(big.NewInt(250)); other.Int64() != 150 || receipts.total.Int64() != 1100 {
		t.Fatalf("other collector cash %v, weekly receipts %v", other, &receipts.total)
	}
}
