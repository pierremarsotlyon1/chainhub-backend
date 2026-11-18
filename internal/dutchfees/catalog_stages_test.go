package dutchfees

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type lotReader func([]call) []result

func (read lotReader) many(_ context.Context, _ *big.Int, qs []call) ([]result, error) {
	return read(qs), nil
}

func TestCatalogStageWindow(t *testing.T) {
	cfg := chainConfig{Burner: common.HexToAddress("0xbd")}
	tokens := map[common.Address]tokenMeta{}
	for i := int64(1); i <= 4; i++ {
		tokens[common.BigToAddress(big.NewInt(i))] = tokenMeta{}
	}
	reader := lotReader(func(qs []call) []result {
		rows := make([]result, len(qs))
		for i, q := range qs {
			a := common.BigToAddress(big.NewInt(int64(i + 1)))
			if q.To != cfg.Burner || !bytes.Equal(q.Data, pack("lots(address)", addrWord(a))) {
				t.Fatalf("bad getter %v", q)
			}
			rows[i] = result{true, append(uintWord(uint64(i+9)), uintWord(uint64(i))...)}
		}
		return rows
	})
	got, err := catalogStages(context.Background(), reader, cfg, big.NewInt(33), tokens, 10, 12)
	if err != nil || len(got) != 2 || got[common.HexToAddress("0x02")].Int64() != 1 || got[common.HexToAddress("0x03")].Int64() != 2 {
		t.Fatal(got, err)
	}
	for _, bad := range []result{{false, make([]byte, 64)}, {true, make([]byte, 32)}, {true, bytes.Repeat([]byte{255}, 64)}} {
		fail := lotReader(func(qs []call) []result { rows := reader(qs); rows[len(rows)-1] = bad; return rows })
		got, err := catalogStages(context.Background(), fail, cfg, big.NewInt(33), tokens, 10, 12)
		if got != nil || err == nil {
			t.Fatal("accepted partial read", bad)
		}
	}
}
func TestHistoryUnavailable(t *testing.T) {
	for _, err := range []error{errors.New("ErrUpstreamBlockUnavailable"), errors.New("block not found for eth_getLogs"), errors.New("getLogs request exceeded max allowed range"), fmt.Errorf("flow logs: %w", rpc.HTTPError{StatusCode: 504})} {
		if !historyUnavailable(err) {
			t.Error(err)
		}
	}
	for _, err := range []error{nil, context.Canceled, fmt.Errorf("flow logs: %w", context.DeadlineExceeded), errors.New("malformed flow event in 0x429 at log 10"), errors.New("conflicting hashes for block 16"), errors.New("invalid or removed log in block 129"), errors.New("invalid API key")} {
		if historyUnavailable(err) {
			t.Error(err)
		}
	}
}
