package dutchfees

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"
	mc "main/contracts/multicall"
)

type call struct {
	To   common.Address
	Data []byte
}
type result struct {
	OK   bool
	Data []byte
}
type snapshot struct {
	Number    uint64      `json:"number"`
	Timestamp uint64      `json:"timestamp"`
	Hash      common.Hash `json:"hash"`
}
type reader struct {
	client   *ethclient.Client
	requests atomic.Uint64
}
type countedHTTP struct{ count *atomic.Uint64 }

func (t countedHTTP) RoundTrip(q *http.Request) (*http.Response, error) {
	t.count.Add(1)
	return http.DefaultTransport.RoundTrip(q)
}

func openReader(ctx context.Context, endpoint string, expectedID uint64) (*reader, error) {
	u, err := url.Parse(endpoint)
	if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") {
		return nil, errors.New("RPC endpoint must be an HTTP(S) URL")
	}
	r := new(reader)
	client, err := rpc.DialOptions(ctx, endpoint, rpc.WithHTTPClient(&http.Client{Transport: countedHTTP{&r.requests}, Timeout: 45 * time.Second}))
	if err != nil {
		return nil, err
	}
	r.client = ethclient.NewClient(client)
	id, err := r.client.ChainID(ctx)
	if err != nil {
		r.client.Close()
		return nil, fmt.Errorf("RPC chain ID: %w", err)
	}
	if !id.IsUint64() || id.Uint64() != expectedID {
		r.client.Close()
		return nil, fmt.Errorf("RPC returned chain %s, expected %d", id, expectedID)
	}
	return r, nil
}
func addrWord(a common.Address) []byte { return common.LeftPadBytes(a.Bytes(), 32) }
func uintWord(n uint64) []byte         { return new(big.Int).SetUint64(n).FillBytes(make([]byte, 32)) }
func pack(sig string, words ...[]byte) []byte {
	out := make([]byte, 4, 4+32*len(words))
	copy(out, crypto.Keccak256([]byte(sig))[:4])
	for _, word := range words {
		out = append(out, word...)
	}
	return out
}
func amountText(amount *big.Int, decimals int) string {
	digits := new(big.Int).Abs(amount).String()
	if decimals > 0 {
		if len(digits) <= decimals {
			digits = strings.Repeat("0", decimals+1-len(digits)) + digits
		}
		cut := len(digits) - decimals
		digits = strings.TrimRight(strings.TrimRight(digits[:cut]+"."+digits[cut:], "0"), ".")
	}
	if amount.Sign() < 0 {
		return "-" + digits
	}
	return digits
}
func (r *reader) many(ctx context.Context, block *big.Int, calls []call) ([]result, error) {
	if block == nil {
		return nil, errors.New("multicall requires a pinned block")
	}
	out := make([]result, len(calls))
	contractABI, err := mc.MulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(3)
	const chunk = 128
	for start := 0; start < len(calls); start += chunk {
		end := min(start+chunk, len(calls))
		group.Go(func() error {
			batch := make([]mc.Multicall3Call, end-start)
			for i, q := range calls[start:end] {
				batch[i] = mc.Multicall3Call{Target: q.To, CallData: q.Data}
			}
			data, err := contractABI.Pack("tryAggregate", false, batch)
			if err != nil {
				return err
			}
			raw, err := r.client.CallContract(ctx, ethereum.CallMsg{To: &poolMulticall, Data: data}, block)
			if err != nil {
				return fmt.Errorf("multicall at block %s: %w", block, err)
			}
			var rows []mc.Multicall3Result
			if err = contractABI.UnpackIntoInterface(&rows, "tryAggregate", raw); err != nil {
				return fmt.Errorf("decode multicall: %w", err)
			}
			if len(rows) != len(batch) {
				return fmt.Errorf("multicall returned %d of %d rows", len(rows), len(batch))
			}
			for i, row := range rows {
				out[start+i] = result{row.Success, row.ReturnData}
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}
	return out, nil
}
func (r *reader) header(ctx context.Context, tag string) (snapshot, error) {
	var h *struct {
		Number    *hexutil.Uint64
		Timestamp *hexutil.Uint64
		Hash      common.Hash
	}
	err := r.client.Client().CallContext(ctx, &h, "eth_getBlockByNumber", tag, false)
	if err != nil {
		return snapshot{}, err
	}
	if h == nil || h.Number == nil || h.Timestamp == nil || h.Hash == (common.Hash{}) {
		return snapshot{}, fmt.Errorf("incomplete RPC block %s", tag)
	}
	n := uint64(*h.Number)
	if tag != "latest" && tag != hexutil.EncodeUint64(n) {
		return snapshot{}, fmt.Errorf("requested block %s, got %d", tag, n)
	}
	return snapshot{Number: n, Timestamp: uint64(*h.Timestamp), Hash: h.Hash}, nil
}
func (r *reader) head(ctx context.Context) (snapshot, error) { return r.header(ctx, "latest") }
func (r *reader) checkHead(ctx context.Context, head snapshot) error {
	h, err := r.header(ctx, hexutil.EncodeUint64(head.Number))
	if err != nil {
		return err
	}
	if h.Hash != head.Hash {
		return fmt.Errorf("block %d reorganized during snapshot", head.Number)
	}
	return nil
}

// Find the earliest block with timestamp >= start, including chains with repeated timestamps.
func (r *reader) blockAtStart(ctx context.Context, head snapshot, start uint64, cfg chainConfig) (uint64, error) {
	if start > head.Timestamp {
		return 0, fmt.Errorf("week starts after pinned head %d", head.Number)
	}
	upper := head
	delta := head.Timestamp - start
	switch strings.ToLower(cfg.Name) {
	case "ethereum":
		delta /= 12
	case "base":
		delta *= 5
	case "arbitrum":
		delta *= 4
	}
	step := max(uint64(1), delta)
	for {
		n := head.Number - min(step, head.Number)
		h, err := r.header(ctx, hexutil.EncodeUint64(n))
		if err != nil {
			return 0, err
		}
		if h.Timestamp < start {
			lower := h
			for probe := 0; upper.Number-lower.Number > 1; probe++ {
				span := upper.Number - lower.Number
				// Linear interpolation converges rapidly on chains with stable block times.
				offset := uint64(float64(span) * float64(start-lower.Timestamp) / float64(upper.Timestamp-lower.Timestamp))
				offset = min(max(offset, uint64(1)), span-1)
				if upper.Timestamp == start {
					offset = span - 1
				}
				if probe >= 3 {
					offset = span / 2
				}
				check, err := r.header(ctx, hexutil.EncodeUint64(lower.Number+offset))
				if err != nil {
					return 0, err
				}
				if check.Timestamp < start {
					lower = check
				} else {
					upper = check
				}
			}
			return upper.Number, nil
		}
		if n == 0 {
			return 0, nil
		}
		upper = h
		if step > head.Number/2 {
			step = head.Number
		} else {
			step *= 2
		}
	}
}
