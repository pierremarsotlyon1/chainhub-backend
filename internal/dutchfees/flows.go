package dutchfees

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type flows struct {
	NativeTakenCrvUSD      string `json:"nativeTakenCrvUSD"`
	CowTradeCrvUSD         string `json:"cowTradeCrvUSD"`
	RealizedCrvUSD         string `json:"realizedCrvUSD"`
	DistributedGrossCrvUSD string `json:"distributedGrossCrvUSD,omitempty"`
	DistributedVeCRVCrvUSD string `json:"distributedVeCRVCrvUSD,omitempty"`
	BridgeInitiatedCrvUSD  string `json:"bridgeInitiatedCrvUSD,omitempty"`
}

var (
	lotStagedEvent       = flowTopic("LotStaged(address,uint256,uint256,uint256,uint256,uint256)")
	nativeTakenEvent     = flowTopic("Taken(address,address,address,uint256,uint256,uint256)")
	cowTradeEvent        = flowTopic("Trade(address,address,address,uint256,uint256,uint256,bytes)")
	feeDistributionEvent = flowTopic("FeesDistributed(uint256,uint256)")
	bridgeInitiatedEvent = flowTopic("Bridge(address,address,address,uint256)")
	cowSettlement        = common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")
)

func flowTopic(signature string) common.Hash { return crypto.Keccak256Hash([]byte(signature)) }
func flowWord(data []byte, i int) *big.Int   { return new(big.Int).SetBytes(data[i*32 : (i+1)*32]) }

func weeklyFlows(ctx context.Context, r *reader, cfg chainConfig, from, to uint64, target, hooker common.Address) (flows, map[common.Address]*big.Int, error) {
	addresses := []common.Address{cfg.Burner}
	signatures := []common.Hash{lotStagedEvent, nativeTakenEvent}
	if cfg.Allocator != (common.Address{}) {
		addresses = append(addresses, cfg.Allocator)
		signatures = append(signatures, feeDistributionEvent)
	}
	if cfg.Bridger != (common.Address{}) {
		addresses = append(addresses, cfg.Bridger)
		signatures = append(signatures, bridgeInitiatedEvent)
	}
	protocolLogs, err := flowLogs(ctx, r.client, ethereum.FilterQuery{Addresses: addresses, Topics: [][]common.Hash{signatures}}, from, to)
	if err != nil {
		return flows{}, nil, err
	}
	// CoW can settle only a lot staged by this burner in this epoch.
	// Installation at the snapshot alone cannot rule out an earlier trade.
	first := to
	hasLot := false
	for _, l := range protocolLogs {
		if l.Address == cfg.Burner && len(l.Topics) > 0 && l.Topics[0] == lotStagedEvent {
			if l.BlockNumber < from || l.BlockNumber > to {
				return flows{}, nil, fmt.Errorf("out-of-range LotStaged block %d", l.BlockNumber)
			}
			hasLot = true
			if l.BlockNumber < first {
				first = l.BlockNumber
			}
		}
	}
	if hasLot {
		cowLogs, err := flowLogs(ctx, r.client, ethereum.FilterQuery{Addresses: []common.Address{cowSettlement}, Topics: [][]common.Hash{{cowTradeEvent}, {common.BytesToHash(cfg.Burner.Bytes())}}}, first, to)
		if err != nil {
			return flows{}, nil, err
		}
		protocolLogs = append(protocolLogs, cowLogs...)
	}
	return summarizeFlowLogs(cfg, target, hooker, protocolLogs)
}

// weeklyStages reads only the proposed burner's staging events in the supplied
// collection interval. The last staging event for each token sets its lot size.
func weeklyStages(ctx context.Context, r *reader, cfg chainConfig, from, to uint64) (map[common.Address]*big.Int, error) {
	logs, err := flowLogs(ctx, r.client, ethereum.FilterQuery{Addresses: []common.Address{cfg.Burner}, Topics: [][]common.Hash{{lotStagedEvent}}}, from, to)
	if err != nil {
		return nil, err
	}
	_, staged, err := summarizeFlowLogs(cfg, common.Address{}, common.Address{}, logs)
	return staged, err
}

func summarizeFlowLogs(cfg chainConfig, target, hooker common.Address, logs []types.Log) (flows, map[common.Address]*big.Int, error) {
	staged := make(map[common.Address]*big.Int)
	native, cow, gross, ve, bridged := new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	sort.Slice(logs, func(i, j int) bool {
		if logs[i].BlockNumber != logs[j].BlockNumber {
			return logs[i].BlockNumber < logs[j].BlockNumber
		}
		return logs[i].Index < logs[j].Index
	})
	var lastBlock uint64
	var lastHash common.Hash
	var lastIndex uint
	for i, l := range logs {
		if l.Removed || len(l.Topics) == 0 {
			return flows{}, nil, fmt.Errorf("invalid or removed log in block %d", l.BlockNumber)
		}
		if i > 0 && lastBlock == l.BlockNumber {
			if lastHash != l.BlockHash {
				return flows{}, nil, fmt.Errorf("conflicting hashes for block %d", l.BlockNumber)
			}
			if lastIndex == l.Index {
				continue
			}
		}
		lastBlock, lastHash, lastIndex = l.BlockNumber, l.BlockHash, l.Index
		malformed := func() error { return fmt.Errorf("malformed flow event in %s at log %d", l.TxHash, l.Index) }
		switch {
		case l.Address == cfg.Burner && l.Topics[0] == lotStagedEvent:
			if len(l.Topics) != 2 || len(l.Data) != 160 {
				return flows{}, nil, malformed()
			}
			staged[common.BytesToAddress(l.Topics[1].Bytes())] = flowWord(l.Data, 0)
		case l.Address == cfg.Burner && l.Topics[0] == nativeTakenEvent:
			if len(l.Topics) != 3 || len(l.Data) != 128 {
				return flows{}, nil, malformed()
			}
			native.Add(native, flowWord(l.Data, 2))
		case l.Address == cowSettlement && l.Topics[0] == cowTradeEvent:
			if len(l.Topics) != 2 || len(l.Data) < 192 {
				return flows{}, nil, malformed()
			}
			if common.BytesToAddress(l.Topics[1].Bytes()) == cfg.Burner && common.BytesToAddress(l.Data[32:64]) == target {
				cow.Add(cow, flowWord(l.Data, 3))
			}
		case l.Address == cfg.Allocator && l.Topics[0] == feeDistributionEvent:
			if len(l.Topics) != 1 || len(l.Data) != 64 {
				return flows{}, nil, malformed()
			}
			gross.Add(gross, flowWord(l.Data, 0))
			ve.Add(ve, flowWord(l.Data, 1))
		case l.Address == cfg.Bridger && l.Topics[0] == bridgeInitiatedEvent:
			if len(l.Topics) != 4 || len(l.Data) != 32 {
				return flows{}, nil, malformed()
			}
			if common.BytesToAddress(l.Topics[1].Bytes()) == target && common.BytesToAddress(l.Topics[2].Bytes()) == hooker && common.BytesToAddress(l.Topics[3].Bytes()) == cfg.Treasury {
				bridged.Add(bridged, flowWord(l.Data, 0))
			}
		}
	}
	out := flows{NativeTakenCrvUSD: amountText(native, 18), CowTradeCrvUSD: amountText(cow, 18), RealizedCrvUSD: amountText(new(big.Int).Add(native, cow), 18)}
	if cfg.Allocator != (common.Address{}) {
		out.DistributedGrossCrvUSD = amountText(gross, 18)
		out.DistributedVeCRVCrvUSD = amountText(ve, 18)
	}
	if cfg.Bridger != (common.Address{}) {
		out.BridgeInitiatedCrvUSD = amountText(bridged, 18)
	}
	return out, staged, nil
}

// Start with the full week and reuse a successful provider range after a size limit.
func flowLogs(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery, from, to uint64) ([]types.Log, error) {
	if from > to {
		return nil, fmt.Errorf("invalid flow block interval %d..%d", from, to)
	}
	size := to - from + 1
	var logs []types.Log
	for first := from; ; {
		last := first + size - 1
		if last < first || last > to {
			last = to
		}
		q.FromBlock = new(big.Int).SetUint64(first)
		q.ToBlock = new(big.Int).SetUint64(last)
		page, err := client.FilterLogs(ctx, q)
		for retry := 0; err != nil && retry < 3 && flowRateLimited(err); retry++ {
			pause := time.NewTimer((time.Second / 2) << retry)
			select {
			case <-ctx.Done():
				pause.Stop()
				return nil, ctx.Err()
			case <-pause.C:
			}
			page, err = client.FilterLogs(ctx, q)
		}
		if err != nil {
			if last > first && flowRangeLimit(err) {
				size = (last - first + 1) / 2
				if limit := flowProviderLimit(err); limit > 0 && limit < last-first+1 {
					size = limit
				}
				continue
			}
			return nil, fmt.Errorf("flow logs %d..%d: %w", first, last, err)
		}
		logs = append(logs, page...)
		if last == to {
			return logs, nil
		}
		first = last + 1
	}
}

func flowRangeLimit(err error) bool {
	msg := strings.ToLower(err.Error())
	for _, s := range []string{"max allowed range", "413 request entity too large", "413 content too large", "block range", "range limit", "maximum range", "range is too", "too many results", "query returned more", "too many logs", "response size", "limit the query", "exceeds max results"} {
		if strings.Contains(msg, s) {
			return true
		}
	}
	return false
}

func flowRateLimited(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "429") || strings.Contains(msg, "too many requests") || strings.Contains(msg, "rate limit")
}

// The CRVHub JSON-RPC gateway advertises its eth_getLogs page size in error.data.
func flowProviderLimit(err error) uint64 {
	var withData interface{ ErrorData() interface{} }
	if !errors.As(err, &withData) {
		return 0
	}
	raw, e := json.Marshal(withData.ErrorData())
	if e != nil {
		return 0
	}
	var data struct {
		Details struct {
			MaxAllowedRange uint64 `json:"maxAllowedRange"`
		} `json:"details"`
	}
	if json.Unmarshal(raw, &data) != nil {
		return 0
	}
	return data.Details.MaxAllowedRange
}
