package dutchfees

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestRouteFromHook(t *testing.T) {
	cfg := chainConfig{Bridger: common.HexToAddress("0x01"), Target: common.HexToAddress("0x02"), Treasury: common.HexToAddress("0x03")}
	makeHook := func(to common.Address, data []byte, duty uint64) []byte {
		var b []byte
		for _, w := range [][]byte{uintWord(32), addrWord(to), uintWord(320), make([]byte, 7*32), uintWord(duty), uintWord(uint64(len(data))), common.RightPadBytes(data, (len(data)+31)/32*32)} {
			b = append(b, w...)
		}
		return b
	}
	calldata := pack("bridge(address,address,uint256,uint256)", addrWord(cfg.Target), addrWord(cfg.Treasury), bytes.Repeat([]byte{255}, 32), uintWord(100))
	hook := makeHook(cfg.Bridger, calldata, 1)
	if ok, min := routeFromHook(hook, cfg); !ok || min.Cmp(big.NewInt(100)) != 0 {
		t.Fatal(ok, min)
	}
	for _, bad := range [][]byte{hook[:len(hook)-1], makeHook(cfg.Bridger, calldata, 0), makeHook(cfg.Target, calldata, 1)} {
		if ok, _ := routeFromHook(bad, cfg); ok {
			t.Fatal("invalid hook verified")
		}
	}
	wrongTreasury := cfg
	wrongTreasury.Treasury = cfg.Target
	if ok, _ := routeFromHook(hook, wrongTreasury); ok {
		t.Fatal("wrong receiver verified")
	}
	mainnet := chainConfig{Allocator: cfg.Treasury}
	if ok, _ := routeFromHook(makeHook(mainnet.Allocator, pack("distribute_fees()"), 1), mainnet); !ok {
		t.Fatal("allocator not verified")
	}
}
