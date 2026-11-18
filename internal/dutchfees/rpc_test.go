package dutchfees

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type flowHeaderStub struct{}

func (flowHeaderStub) GetBlockByNumber(tag string, full bool) (map[string]any, error) {
	n := uint64(980)
	if tag != "latest" {
		v, err := hexutil.DecodeUint64(tag)
		if err != nil {
			return nil, err
		}
		n = v
	}
	return map[string]any{"number": hexutil.EncodeUint64(n), "timestamp": hexutil.EncodeUint64(200 + n/10), "hash": common.BigToHash(new(big.Int).SetUint64(1 + n))}, nil
}
func TestRPCWeekLowerBound(t *testing.T) {
	server := rpc.NewServer()
	defer server.Stop()
	if err := server.RegisterName("eth", flowHeaderStub{}); err != nil {
		t.Fatal(err)
	}
	r := &reader{client: ethclient.NewClient(rpc.DialInProc(server))}
	defer r.client.Close()
	ctx := context.Background()
	h, err := r.head(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for ts, want := range map[uint64]uint64{190: 0, 200: 0, 201: 10, 250: 500, 298: 980} {
		got, err := r.blockAtStart(ctx, h, ts, chainConfig{Name: "arbitrum"})
		if err != nil || got != want {
			t.Fatalf("at %d, got %d, want %d, err %v", ts, got, want, err)
		}
	}
	if err = r.checkHead(ctx, h); err != nil {
		t.Fatal(err)
	}
	h.Hash[0] ^= 1
	if r.checkHead(ctx, h) == nil {
		t.Fatal("accepted reorganized head")
	}
}
func TestAmountTextPreservesPrecision(t *testing.T) {
	for _, tt := range []struct {
		v int64
		d int
		s string
	}{{0, 18, "0"}, {10, 0, "10"}, {1000, 2, "10"}, {-1, 18, "-0.000000000000000001"}, {1050, 2, "10.5"}} {
		if got := amountText(big.NewInt(tt.v), tt.d); got != tt.s {
			t.Fatalf("got %s, want %s", got, tt.s)
		}
	}
}
