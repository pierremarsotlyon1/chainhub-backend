package dutchfees

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"os"
	"testing"

	"main/utils"
)

type splitterFixture struct {
	t         *testing.T
	defaultOK bool
	failProbe bool
}

func (f splitterFixture) many(_ context.Context, block *big.Int, calls []call) ([]result, error) {
	f.t.Helper()
	if block.Cmp(big.NewInt(123)) != 0 {
		f.t.Errorf("wrong block: %s", block)
	}
	cfg := networks[0]
	splitter := utils.CRV_USD_FEE_SPLITTER
	good1, bad, good2 := uintWord(10), uintWord(11), uintWord(12)
	if len(calls) == 3 && calls[0].To == cfg.Target {
		bal := pack("balanceOf(address)", addrWord(cfg.Collector))
		if !bytes.Equal(calls[0].Data, bal) || !bytes.Equal(calls[2].Data, bal) || calls[2].To != cfg.Target || calls[1].To != splitter {
			f.t.Error("unexpected measurement order or recipient")
		}
		ok := f.defaultOK
		if !bytes.Equal(calls[1].Data, pack("dispatch_fees()")) {
			if !bytes.Equal(calls[1].Data, pack("dispatch_fees(address[])", uintWord(32), uintWord(2), good1, good2)) {
				f.t.Error("wrong subset in final dispatch")
			}
			ok = true
		}
		// Budget: 100 already held by collector; 20 idle on splitter, and 3+4 newly claimable.
		after := uint64(100)
		if ok {
			after += 27
		}
		return []result{{true, uintWord(100)}, {ok, nil}, {true, uintWord(after)}}, nil
	}
	if len(calls) == 1 && bytes.Equal(calls[0].Data, pack("n_controllers()")) {
		return []result{{true, uintWord(3)}}, nil
	}
	if len(calls) == 3 {
		for i, c := range calls {
			if c.To != splitter || !bytes.Equal(c.Data, pack("controllers(uint256)", uintWord(uint64(i)))) {
				f.t.Error("bad controller enumeration")
			}
		}
		return []result{{true, good1}, {true, bad}, {true, good2}}, nil
	}
	if len(calls) != 1 || calls[0].To != splitter {
		return nil, errors.New("unexpected probe")
	}
	if f.failProbe {
		return nil, errors.New("RPC offline")
	}
	for _, v := range [][]byte{good1, bad, good2} {
		if bytes.Equal(calls[0].Data, pack("dispatch_fees(address[])", uintWord(32), uintWord(1), v)) {
			return []result{{!bytes.Equal(v, bad), nil}}, nil
		}
	}
	return nil, errors.New("unexpected dispatch")
}

func TestControllerReceipts(t *testing.T) {
	for _, defaultOK := range []bool{true, false} {
		amount, partial, err := controllerReceipts(context.Background(), splitterFixture{t: t, defaultOK: defaultOK}, networks[0], big.NewInt(123))
		if err != nil || amount.Cmp(big.NewInt(27)) != 0 || partial == defaultOK {
			t.Fatalf("default success %t: amount %v partial %t err %v", defaultOK, amount, partial, err)
		}
	}
	if _, _, err := controllerReceipts(context.Background(), splitterFixture{t: t, failProbe: true}, networks[0], big.NewInt(123)); err == nil {
		t.Error("RPC failure mistaken for an unclaimable controller")
	}
}

func TestControllerReceiptsLive(t *testing.T) {
	endpoint := os.Getenv("DUTCHFEES_TEST_ETH_RPC")
	if endpoint == "" {
		t.Skip("set DUTCHFEES_TEST_ETH_RPC for the live Ethereum smoke test")
	}
	ctx := context.Background()
	r, err := openReader(ctx, endpoint, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer r.client.Close()
	head, err := r.head(ctx)
	if err != nil {
		t.Fatal(err)
	}
	amount, partial, err := controllerReceipts(ctx, r, networks[0], new(big.Int).SetUint64(head.Number))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("block %d controller receipts %s crvUSD partial %t", head.Number, amountText(amount, 18), partial)
}
