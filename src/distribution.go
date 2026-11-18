package src

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"main/utils"
)

const allocatorFirstBlock uint64 = 19526061
const distributionWeek uint64 = 7 * 86400

func fetchLastDistribution(ctx context.Context, client *ethclient.Client, head uint64) (uint64, uint64, error) {
	topic := crypto.Keccak256Hash([]byte("FeesDistributed(uint256,uint256)"))
	logs := func(from, to uint64) ([]types.Log, error) {
		return legacyTransferLogs(ctx, client, ethereum.FilterQuery{FromBlock: new(big.Int).SetUint64(from), ToBlock: new(big.Int).SetUint64(to), Addresses: []common.Address{utils.FEE_ALLOCATOR}, Topics: [][]common.Hash{{topic}}})
	}
	timestamp := func(block uint64) (uint64, error) {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(block))
		if err != nil {
			return 0, err
		}
		return h.Time, nil
	}
	return latestDistributionWeek(head, logs, timestamp)
}

// Rebuild the latest nonempty Thursday UTC week, including on upgrade from snapshots
// that stored the gross allocator amount. Disjoint ranges keep retries idempotent.
func latestDistributionWeek(head uint64, logs func(uint64, uint64) ([]types.Log, error), timestamp func(uint64) (uint64, error)) (uint64, uint64, error) {
	var latest uint64
	total := new(big.Int)
	span := uint64(8192)
	for end := head; end >= allocatorFirstBlock; {
		start := end - min(end-allocatorFirstBlock, span-1)
		events, err := logs(start, end)
		if err != nil {
			return 0, 0, err
		}
		seen := map[[2]uint64]bool{}
		times := map[uint64]uint64{}
		for _, event := range events {
			if event.Removed || event.BlockNumber < start || event.BlockNumber > end || len(event.Data) != 64 {
				return 0, 0, fmt.Errorf("invalid allocator log at block %d", event.BlockNumber)
			}
			key := [2]uint64{event.BlockNumber, uint64(event.Index)}
			if seen[key] {
				continue
			}
			seen[key] = true
			ts, ok := times[event.BlockNumber]
			if !ok {
				ts, err = timestamp(event.BlockNumber)
				if err != nil {
					return 0, 0, err
				}
				times[event.BlockNumber] = ts
			}
			if ts/distributionWeek > latest/distributionWeek {
				total.SetUint64(0)
			}
			if ts/distributionWeek >= latest/distributionWeek {
				total.Add(total, new(big.Int).SetBytes(event.Data[32:]))
				latest = max(latest, ts)
			}
		}
		if latest != 0 {
			earliest, err := timestamp(start)
			if err != nil {
				return 0, 0, err
			}
			if earliest <= latest-latest%distributionWeek {
				break
			}
		}
		// Widen the search through quiet periods. The RPC reader splits
		// queries when a provider restricts their size.
		if latest == 0 {
			span = min(span*2, 262144)
		}
		end = start - 1
	}
	whole := total.Quo(total, big.NewInt(1e18))
	if !whole.IsUint64() {
		return 0, 0, fmt.Errorf("veCRV weekly distribution exceeds uint64")
	}
	return latest, whole.Uint64(), nil
}
