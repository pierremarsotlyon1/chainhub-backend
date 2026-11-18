package dutchfees

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type protocol struct {
	Installed, Target, Hooker                common.Address
	Phase                                    string
	GlobalMask, TargetMask                   uint64
	CollectorBuffer, HookCash, AllocatorCash *big.Int
	Weight                                   uint64
	RouteVerified                            bool
	CollectStart, CollectEnd, ExchangeEnd    uint64
	ForwardRate, BridgeMinimum               *big.Int
}

func routeFromHook(data []byte, cfg chainConfig) (bool, *big.Int) {
	// Hooker.hooks(uint256) returns one dynamic tuple. Its head contains the
	// destination, bytes offset, seven inlined compensation words and duty flag.
	if len(data) < 32*12 || !bytes.Equal(data[:32], uintWord(32)) {
		return false, nil
	}
	body := data[32:]
	if !bytes.Equal(body[32:64], uintWord(10*32)) || !bytes.Equal(body[9*32:10*32], uintWord(1)) {
		return false, nil
	}
	n := new(big.Int).SetBytes(body[10*32 : 11*32])
	if !n.IsUint64() || n.Uint64() > 1024 {
		return false, nil
	}
	size := int(n.Uint64())
	if len(body) != 11*32+32*((size+31)/32) {
		return false, nil
	}
	foreplay := body[11*32 : 11*32+size]
	if cfg.Allocator != (common.Address{}) {
		return bytes.Equal(body[:32], addrWord(cfg.Allocator)) && bytes.Equal(foreplay, pack("distribute_fees()")), nil
	}
	if cfg.Bridger == (common.Address{}) || cfg.Treasury == (common.Address{}) || !bytes.Equal(body[:32], addrWord(cfg.Bridger)) || len(foreplay) != 132 {
		return false, nil
	}
	// Both wrappers accept MAX_UINT256 to bridge all available target tokens.
	prefix := pack("bridge(address,address,uint256,uint256)", addrWord(cfg.Target), addrWord(cfg.Treasury), bytes.Repeat([]byte{255}, 32))
	if !bytes.Equal(foreplay[:100], prefix) {
		return false, nil
	}
	return true, new(big.Int).SetBytes(foreplay[100:])
}

func readProtocol(ctx context.Context, r *reader, cfg chainConfig, head snapshot) (protocol, error) {
	p := protocol{AllocatorCash: new(big.Int), BridgeMinimum: new(big.Int)}
	block := new(big.Int).SetUint64(head.Number)
	c := cfg.Collector
	calls := []call{
		{c, pack("burner()")}, {c, pack("target()")}, {c, pack("hooker()")}, {c, pack("epoch()")},
		{c, pack("is_killed(address)", addrWord(common.Address{}))}, {c, pack("is_killed(address)", addrWord(cfg.Target))},
	}
	for _, phase := range []uint64{2, 4, 8} {
		calls = append(calls, call{c, pack("epoch_time_frame(uint256,uint256)", uintWord(phase), uintWord(head.Timestamp))})
	}
	rows, err := r.many(ctx, block, calls)
	if err != nil {
		return p, err
	}
	if len(rows) != len(calls) {
		return p, fmt.Errorf("collector getter count mismatch")
	}
	for i, row := range rows {
		size := 32
		if i >= 6 {
			size = 64
		}
		if !row.OK || len(row.Data) != size {
			return p, fmt.Errorf("collector %s getter %x returned invalid data", c, calls[i].Data[:4])
		}
	}
	for i, addr := range []*common.Address{&p.Installed, &p.Target, &p.Hooker} {
		if !bytes.Equal(rows[i].Data[:12], make([]byte, 12)) {
			return p, fmt.Errorf("invalid protocol address")
		}
		*addr = common.BytesToAddress(rows[i].Data)
	}
	if p.Target != cfg.Target {
		return p, fmt.Errorf("collector target %s differs from expected crvUSD %s", p.Target, cfg.Target)
	}
	small := func(word []byte) (uint64, bool) { v := new(big.Int).SetBytes(word); return v.Uint64(), v.IsUint64() }
	phase, ok := small(rows[3].Data)
	p.Phase = map[uint64]string{1: "sleep", 2: "collect", 4: "exchange", 8: "forward"}[phase]
	if !ok || p.Phase == "" {
		return p, fmt.Errorf("unknown epoch %x", rows[3].Data)
	}
	p.GlobalMask, ok = small(rows[4].Data)
	if !ok {
		return p, fmt.Errorf("invalid global kill mask")
	}
	p.TargetMask, ok = small(rows[5].Data)
	if !ok {
		return p, fmt.Errorf("invalid target kill mask")
	}
	var starts, ends [3]uint64
	for i := range starts {
		start, ok1 := small(rows[6+i].Data[:32])
		end, ok2 := small(rows[6+i].Data[32:])
		if !ok1 || !ok2 || start >= end || (i > 0 && start != ends[i-1]) {
			return p, fmt.Errorf("invalid collector epoch time frame")
		}
		starts[i], ends[i] = start, end
	}
	p.CollectStart, p.CollectEnd, p.ExchangeEnd = starts[0], ends[0], ends[1]
	feeTime := starts[2]
	if p.Phase == "forward" {
		feeTime = head.Timestamp
	}
	calls = []call{{c, pack("fee(uint256,uint256)", uintWord(8), uintWord(feeTime))}, {p.Hooker, pack("buffer_amount()")}, {cfg.Target, pack("balanceOf(address)", addrWord(p.Hooker))}, {p.Hooker, pack("hooks(uint256)", uintWord(0))}}
	if cfg.Allocator != (common.Address{}) {
		calls = append(calls, call{cfg.Allocator, pack("distributor_weight()")}, call{cfg.Target, pack("balanceOf(address)", addrWord(cfg.Allocator))})
	}
	dutchInstalled := cfg.Burner != (common.Address{}) && p.Installed == cfg.Burner
	if dutchInstalled {
		calls = append(calls, call{cfg.Burner, pack("want()")})
	}
	rows, err = r.many(ctx, block, calls)
	if err != nil {
		return p, err
	}
	if len(rows) != len(calls) {
		return p, fmt.Errorf("routing getter count mismatch")
	}
	for i, row := range rows {
		// An unset or changed duty hook disables this route without hiding balances.
		if i != 3 && (!row.OK || len(row.Data) != 32) {
			return p, fmt.Errorf("routing getter %s %x returned invalid data", calls[i].To, calls[i].Data[:4])
		}
	}
	if dutchInstalled && !bytes.Equal(rows[len(rows)-1].Data, addrWord(cfg.Target)) {
		return p, fmt.Errorf("installed Dutch burner %s payment token differs from collector crvUSD %s", cfg.Burner, cfg.Target)
	}
	p.ForwardRate = new(big.Int).SetBytes(rows[0].Data)
	if p.ForwardRate.Cmp(big.NewInt(1e18)) > 0 {
		return p, fmt.Errorf("invalid FORWARD fee")
	}
	p.CollectorBuffer = new(big.Int).SetBytes(rows[1].Data)
	p.HookCash = new(big.Int).SetBytes(rows[2].Data)
	if rows[3].OK {
		verified, minimum := routeFromHook(rows[3].Data, cfg)
		p.RouteVerified = verified
		if minimum != nil {
			p.BridgeMinimum = minimum
		}
	}
	if cfg.Allocator != (common.Address{}) {
		p.Weight, ok = small(rows[4].Data)
		if !ok || p.Weight > 10000 {
			return p, fmt.Errorf("invalid allocator weight")
		}
		p.AllocatorCash.SetBytes(rows[5].Data)
	}
	return p, nil
}
