package dutchfees

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestQuoteUsesRemainingPayment(t *testing.T) {
	a, b, burner := common.HexToAddress("0x01"), common.HexToAddress("0x02"), common.HexToAddress("0x03")
	staged := map[common.Address]*big.Int{a: big.NewInt(300), b: big.NewInt(500)}
	stub := stockStub{}
	key := func(token common.Address) string {
		return burner.Hex() + string(pack("getAmountNeeded(address)", addrWord(token)))
	}
	stub[key(a)] = 17
	stub[key(b)] = 0
	quote, err := auctionQuote(context.Background(), stub, big.NewInt(1), burner, staged)
	if err != nil || quote != "0.000000000000000017" {
		t.Fatalf("got %q, %v", quote, err)
	}
}
