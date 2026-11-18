package dutchfees

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/sync/errgroup"
	"main/utils"
)

// controllerReceipts measures only new crvUSD transferred to the Ethereum FeeCollector.
func controllerReceipts(ctx context.Context, r stockReader, cfg chainConfig, block *big.Int) (*big.Int, bool, error) {
	if cfg.ID != 1 || block == nil {
		return nil, false, fmt.Errorf("controller receipts require a pinned Ethereum block")
	}
	splitter := utils.CRV_USD_FEE_SPLITTER
	balance := call{cfg.Target, pack("balanceOf(address)", addrWord(cfg.Collector))}
	measure := func(dispatch []byte) (*big.Int, bool, error) {
		rows, err := r.many(ctx, block, []call{balance, {splitter, dispatch}, balance})
		if err != nil {
			return nil, false, err
		}
		if len(rows) != 3 || !rows[0].OK || !rows[2].OK || len(rows[0].Data) != 32 || len(rows[2].Data) != 32 {
			return nil, false, fmt.Errorf("invalid collector balance during dispatch")
		}
		if !rows[1].OK {
			return nil, false, nil
		}
		delta := new(big.Int).Sub(new(big.Int).SetBytes(rows[2].Data), new(big.Int).SetBytes(rows[0].Data))
		if delta.Sign() < 0 {
			return nil, false, fmt.Errorf("collector balance decreased during dispatch")
		}
		return delta, true, nil
	}

	amount, ok, err := measure(pack("dispatch_fees()"))
	if err != nil || ok {
		return amount, false, err
	}

	// The splitter claims all registered controllers by default. A stale controller
	// can revert the whole dispatch, even when others can still pay the collector.
	count, err := r.many(ctx, block, []call{{splitter, pack("n_controllers()")}})
	if err != nil {
		return nil, false, err
	}
	if len(count) != 1 || !count[0].OK || len(count[0].Data) != 32 {
		return nil, false, fmt.Errorf("invalid FeeSplitter controller count")
	}
	n := new(big.Int).SetBytes(count[0].Data)
	if !n.IsUint64() || n.Uint64() > 50 {
		return nil, false, fmt.Errorf("controller count exceeds FeeSplitter limit")
	}
	queries := make([]call, int(n.Uint64()))
	for i := range queries {
		queries[i] = call{splitter, pack("controllers(uint256)", uintWord(uint64(i)))}
	}
	rows, err := r.many(ctx, block, queries)
	if err != nil {
		return nil, false, err
	}
	if len(rows) != len(queries) {
		return nil, false, fmt.Errorf("incomplete FeeSplitter controller list")
	}

	eligible := make([]bool, len(rows))
	group, probeCtx := errgroup.WithContext(ctx)
	group.SetLimit(3)
	for i, row := range rows {
		if !row.OK || len(row.Data) != 32 || new(big.Int).SetBytes(row.Data).BitLen() > 160 || common.BytesToAddress(row.Data) == (common.Address{}) {
			return nil, false, fmt.Errorf("invalid FeeSplitter controller %d", i)
		}
	}
	for i, row := range rows {
		group.Go(func() error {
			// Probe each candidate in an independent eth_call at the original block.
			probe, err := r.many(probeCtx, block, []call{{splitter, dispatchControllers([][]byte{row.Data})}})
			if err != nil {
				return err
			}
			if len(probe) != 1 {
				return fmt.Errorf("invalid controller dispatch probe")
			}
			eligible[i] = probe[0].OK
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, false, err
	}
	var selected [][]byte
	for i, ok := range eligible {
		if ok {
			selected = append(selected, rows[i].Data)
		}
	}
	if len(selected) == 0 {
		return new(big.Int), true, nil
	}

	// Only this combined dispatch contributes a value: each probe also redistributes
	// the same idle splitter balance, so summing individual receipts would duplicate it.
	amount, ok, err = measure(dispatchControllers(selected))
	if err != nil {
		return nil, false, err
	}
	if !ok {
		return new(big.Int), true, nil
	}
	return amount, len(selected) != len(rows), nil
}

func dispatchControllers(controllers [][]byte) []byte {
	words := [][]byte{uintWord(32), uintWord(uint64(len(controllers)))}
	return pack("dispatch_fees(address[])", append(words, controllers...)...)
}
