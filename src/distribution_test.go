package src

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestLatestDistributionWeek(t *testing.T) {
	base := allocatorFirstBlock
	head := base + 8192
	week := distributionWeek * 3000
	event := func(block uint64, index uint) types.Log {
		data := make([]byte, 64)
		big.NewInt(9e18).FillBytes(data[:32])
		big.NewInt(75e16).FillBytes(data[32:])
		return types.Log{BlockNumber: block, Index: index, Data: data}
	}
	logs := func(from, to uint64) ([]types.Log, error) {
		if to == head {
			return []types.Log{event(head, 0), event(head, 0)}, nil
		}
		return []types.Log{event(base, 0), event(base, 1)}, nil
	}
	timestamp := func(block uint64) (uint64, error) { return week + (block-base)*12, nil }
	for _, end := range []uint64{head, head} {
		ts, amount, err := latestDistributionWeek(end, logs, timestamp)
		if err != nil || ts != week+8192*12 || amount != 2 {
			t.Fatalf("got %d, %d, %v", ts, amount, err)
		}
	}
	timestamp = func(block uint64) (uint64, error) { return week - 1 + (block-base)*12, nil }
	_, amount, err := latestDistributionWeek(head, logs, timestamp)
	if err != nil || amount != 0 {
		t.Fatalf("previous week included: %d, %v", amount, err)
	}
}
func TestLatestDistributionWeekReadFailure(t *testing.T) {
	unavailable := errors.New("unavailable")
	logs := func(uint64, uint64) ([]types.Log, error) { return nil, unavailable }
	_, _, err := latestDistributionWeek(allocatorFirstBlock, logs, nil)
	if !errors.Is(err, unavailable) {
		t.Fatal(err)
	}
}
