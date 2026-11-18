package dutchfees

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// catalogStages returns current lots for known tokens. A deployment can stage
// assets absent from the catalogue, so callers must mark this discovery partial.
func catalogStages(ctx context.Context, r stockReader, cfg chainConfig, block *big.Int, tokens map[common.Address]tokenMeta, collectStart, collectEnd uint64) (map[common.Address]*big.Int, error) {
	if block == nil || block.Sign() < 0 || collectStart >= collectEnd || len(tokens) == 0 {
		return nil, errors.New("catalogue lot discovery requires a pinned block, a valid collection interval and a nonempty catalogue")
	}
	assets := make([]common.Address, 0, len(tokens))
	for a := range tokens {
		assets = append(assets, a)
	}
	sort.Slice(assets, func(i, j int) bool { return bytes.Compare(assets[i][:], assets[j][:]) < 0 })
	calls := make([]call, len(assets))
	for i, a := range assets {
		calls[i] = call{cfg.Burner, pack("lots(address)", addrWord(a))}
	}
	rows, err := r.many(ctx, block, calls)
	if err != nil {
		return nil, fmt.Errorf("catalogue lot discovery: %w", err)
	}
	if len(rows) != len(assets) {
		return nil, fmt.Errorf("catalogue lot discovery returned %d of %d rows", len(rows), len(assets))
	}
	staged := make(map[common.Address]*big.Int)
	for i, row := range rows {
		if !row.OK || len(row.Data) != 64 {
			return nil, fmt.Errorf("unreadable burner lot for %s", assets[i])
		}
		ts := new(big.Int).SetBytes(row.Data[:32])
		if !ts.IsUint64() {
			return nil, fmt.Errorf("invalid burner lot timestamp for %s", assets[i])
		}
		if ts.Sign() > 0 && ts.Uint64() >= collectStart && ts.Uint64() < collectEnd {
			staged[assets[i]] = new(big.Int).SetBytes(row.Data[32:])
		}
	}
	return staged, nil
}

func historyUnavailable(err error) bool {
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var provider rpc.Error
	if errors.As(err, &provider) && provider.ErrorCode() == -32601 {
		return true
	}
	var httpErr rpc.HTTPError
	if errors.As(err, &httpErr) {
		switch httpErr.StatusCode {
		case 408, 413, 429, 502, 503, 504:
			return true
		}
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	if flowRangeLimit(err) {
		return true
	}
	msg := strings.ToLower(err.Error())
	for _, phrase := range []string{"too many requests", "rate limit", "errupstreamblockunavailable", "block not found for eth_getlogs", "historical logs unavailable", "historical data is not available", "history has been pruned", "pruned history", "method not found", "gateway timeout", "gateway time-out", "upstream request timeout"} {
		if strings.Contains(msg, phrase) {
			return true
		}
	}
	return false
}
