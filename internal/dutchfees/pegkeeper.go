package dutchfees

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Removed PegKeepers may still hold withdrawable profit. Include historical
// deployments from the dashboard and pegkeeper dataset alongside the live registry.
var historicalPegKeepers = []common.Address{
	address("0xaa346781ddd7009caa644a4980f044c50cd2ae22"),
	address("0xe7cd2b4eb1d98cd6a4a48b6071d46401ac7dc5c8"),
	address("0x6b765d07cf966c745b340adca67749fe75b5c345"),
	address("0x1ef89ed0edd93d1ec09e4c07373f69c49f4dccae"),
	address("0x68e31e1edd641b13caeab1ac1be661b19cc021ca"),
	address("0x0b502e48e950095d93e8b739ad146c72b4f6c820"),
	address("0x5b49b9add1ecfe53e19cc2cfc8a33127cd6ba4c6"),
	address("0xff78468340ee322ed63c432bf74d817742b392bf"),
	address("0x9201da0d97caaaff53f01b2fb56767c7072de340"),
	address("0xfb726f57d251ab5c731e5c64ed4f5f94351ef9f3"),
	address("0x3fa20eaa107de08b38a8734063d605d5842fe09c"),
	address("0x0a05ff644878b908ef8eb29542aa88c07d9797d3"),
	address("0x503e1bf274e7a6c64152395ae8eb57ec391f91f8"),
	address("0x338Cb2D827112d989A861cDe87CD9FfD913A1f9D"),
}
var pegKeeperRegulator = address("0x36a04CAffc681fa179558B2Aaba30395CDdd855f")

func pegKeeperReceiptsUSD(ctx context.Context, r stockReader, cfg chainConfig, block *big.Int, tokens map[common.Address]tokenMeta) (float64, bool, error) {
	// The deployed regulator's peg_keepers array has a maximum length of eight.
	registry := make([]call, 8)
	for i := range registry {
		registry[i] = call{pegKeeperRegulator, pack("peg_keepers(uint256)", uintWord(uint64(i)))}
	}
	rows, err := r.many(ctx, block, registry)
	if err != nil {
		return 0, true, err
	}
	keepers := append([]common.Address(nil), historicalPegKeepers...)
	partial := !rows[0].OK
	for _, row := range rows {
		if !row.OK {
			continue
		}
		if len(row.Data) != 128 {
			partial = true
			continue
		}
		keeper := common.BytesToAddress(row.Data[:32])
		if keeper == (common.Address{}) {
			partial = true
			continue
		}
		keepers = append(keepers, keeper)
	}
	seen := map[common.Address]bool{}
	var pools []call
	for _, keeper := range keepers {
		if !seen[keeper] {
			pools = append(pools, call{keeper, pack("pool()")})
			seen[keeper] = true
		}
	}
	rows, err = r.many(ctx, block, pools)
	if err != nil {
		return 0, true, err
	}
	var assets []common.Address
	var probes []call
	for i, row := range rows {
		if !row.OK || len(row.Data) != 32 || common.BytesToAddress(row.Data) == (common.Address{}) {
			partial = true
			continue
		}
		pool := common.BytesToAddress(row.Data)
		assets = append(assets, pool)
		balance := call{pool, pack("balanceOf(address)", addrWord(cfg.Collector))}
		// Four calls per keeper keep each before/after pair inside a single
		// tryAggregate (reader.many batches 128 calls). eth_call discards writes.
		probes = append(probes, balance, call{pools[i].To, pack("withdraw_profit()")}, balance, call{cfg.Collector, pack("is_killed(address)", addrWord(pool))})
	}
	rows, err = r.many(ctx, block, probes)
	if err != nil {
		return 0, true, err
	}
	total := 0.
	for i, asset := range assets {
		usd, known := pegKeeperValue(rows[i*4:i*4+4], tokens[asset])
		total += usd
		partial = partial || !known
	}
	return total, partial, nil
}

func pegKeeperValue(rows []result, token tokenMeta) (float64, bool) {
	for _, row := range rows[:3] {
		if !row.OK || len(row.Data) != 32 {
			return 0, false
		}
	}
	received := new(big.Int).Sub(new(big.Int).SetBytes(rows[2].Data), new(big.Int).SetBytes(rows[0].Data))
	if received.Sign() < 0 {
		return 0, false
	}
	if received.Sign() == 0 {
		return 0, true
	}
	if !rows[3].OK || len(rows[3].Data) != 32 {
		return 0, false
	}
	mask := new(big.Int).SetBytes(rows[3].Data)
	if mask.Bit(1) != 0 || mask.Bit(2) != 0 {
		return 0, true
	}
	return valueUSD(received, token)
}
