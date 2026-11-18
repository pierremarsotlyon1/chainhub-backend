package dutchfees

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

// auctionQuote sums the current payment for all remaining available inventory
// staged this week. The burner returns zero for an inactive or empty lot.
func auctionQuote(ctx context.Context, r stockReader, block *big.Int, burner common.Address, staged map[common.Address]*big.Int) (string, error) {
	if len(staged) == 0 {
		return "0", nil
	}
	tokens := make([]common.Address, 0, len(staged))
	for token := range staged {
		tokens = append(tokens, token)
	}
	sort.Slice(tokens, func(i, j int) bool { return bytes.Compare(tokens[i][:], tokens[j][:]) < 0 })
	calls := make([]call, len(tokens))
	for i, token := range tokens {
		calls[i] = call{burner, pack("getAmountNeeded(address)", addrWord(token))}
	}
	rows, err := r.many(ctx, block, calls)
	if err != nil {
		return "", fmt.Errorf("auction quote: %w", err)
	}
	if len(rows) != len(tokens) {
		return "", fmt.Errorf("auction quote returned %d of %d token results", len(rows), len(tokens))
	}
	total := new(big.Int)
	for i, row := range rows {
		if !row.OK || len(row.Data) != 32 {
			return "", fmt.Errorf("unreadable auction quote for %s at block %s", tokens[i], block)
		}
		total.Add(total, new(big.Int).SetBytes(row.Data))
	}
	return amountText(total, 18), nil
}
