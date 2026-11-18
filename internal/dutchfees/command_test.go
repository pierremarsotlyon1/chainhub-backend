package dutchfees

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	mc "main/contracts/multicall"
	"main/interfaces"
)

// Serve actual JSON-RPC and Multicall3 ABI envelopes. Every read must match
// an expected contract, selector, arguments, and block.
type commandRPC struct {
	cfg     chainConfig
	head    uint64
	replies map[string][]byte
	lot     common.Address
}

func (s *commandRPC) ChainId() hexutil.Uint64 { return hexutil.Uint64(s.cfg.ID) }
func (s *commandRPC) GetBlockByNumber(tag string, _ bool) (map[string]any, error) {
	n := s.head
	if tag != "latest" {
		var err error
		n, err = hexutil.DecodeUint64(tag)
		if err != nil {
			return nil, err
		}
	}
	return map[string]any{"number": hexutil.EncodeUint64(n), "timestamp": hexutil.EncodeUint64(weekSeconds + n*12), "hash": common.BigToHash(new(big.Int).SetUint64(n + 1))}, nil
}
func (s *commandRPC) Call(tx struct {
	To    common.Address
	Input hexutil.Bytes
}, tag string) (hexutil.Bytes, error) {
	contract, err := mc.MulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	m := contract.Methods["tryAggregate"]
	if tag != hexutil.EncodeUint64(s.head) || tx.To != poolMulticall || len(tx.Input) < 4 || !bytes.Equal(tx.Input[:4], m.ID) {
		return nil, fmt.Errorf("unexpected call at %s", tag)
	}
	args, err := m.Inputs.Unpack(tx.Input[4:])
	if err != nil {
		return nil, err
	}
	calls := *abi.ConvertType(args[1], new([]mc.Multicall3Call)).(*[]mc.Multicall3Call)
	rows := make([]mc.Multicall3Result, len(calls))
	for i, c := range calls {
		data, ok := s.replies[c.Target.Hex()+string(c.CallData)]
		if !ok {
			return nil, fmt.Errorf("unexpected call %s %x", c.Target, c.CallData)
		}
		rows[i] = mc.Multicall3Result{Success: true, ReturnData: data}
	}
	return m.Outputs.Pack(rows)
}
func (s *commandRPC) GetLogs(q struct {
	FromBlock, ToBlock string
	Address            []common.Address
	Topics             [][]common.Hash
}) ([]types.Log, error) {
	if s.head != 5 || q.FromBlock != "0x2" || q.ToBlock != "0x3" || len(q.Address) != 1 || q.Address[0] != s.cfg.Burner || len(q.Topics) != 1 || len(q.Topics[0]) != 1 || q.Topics[0][0] != lotStagedEvent {
		return nil, fmt.Errorf("unexpected log filter %+v", q)
	}
	data := make([]byte, 160)
	copy(data[:32], uintWord(10))
	return []types.Log{{Address: s.cfg.Burner, Topics: []common.Hash{lotStagedEvent, common.BytesToHash(s.lot.Bytes())}, Data: data, BlockNumber: 3}}, nil
}

func TestReadChainInstalled(t *testing.T) {
	cfg := networks[0]
	hook, lot := common.HexToAddress("0x10"), common.HexToAddress("0x20")
	wad := func(n int64) []byte {
		return new(big.Int).Mul(big.NewInt(n), big.NewInt(1e18)).FillBytes(make([]byte, 32))
	}
	for _, tt := range []struct {
		phase       string
		epoch, head uint64
		expected    float64
	}{{"forward", 8, 7, 100}, {"exchange", 4, 5, 130}} {
		t.Run(tt.phase, func(t *testing.T) {
			s := &commandRPC{cfg: cfg, head: tt.head, lot: lot, replies: map[string][]byte{}}
			put := func(a common.Address, sig string, value []byte, args ...[]byte) {
				s.replies[a.Hex()+string(pack(sig, args...))] = value
			}
			put(cfg.Collector, "burner()", addrWord(cfg.Burner))
			put(cfg.Collector, "target()", addrWord(cfg.Target))
			put(cfg.Collector, "hooker()", addrWord(hook))
			put(cfg.Collector, "epoch()", uintWord(tt.epoch))
			for _, a := range []common.Address{{}, cfg.Target, lot} {
				put(cfg.Collector, "is_killed(address)", uintWord(0), addrWord(a))
			}
			ts := weekSeconds + tt.head*12
			for i, epoch := range []uint64{2, 4, 8} {
				start := weekSeconds + uint64(i+1)*24
				put(cfg.Collector, "epoch_time_frame(uint256,uint256)", append(uintWord(start), uintWord(start+24)...), uintWord(epoch), uintWord(ts))
			}
			feeTime := weekSeconds + 72
			if tt.phase == "forward" {
				feeTime = ts
			}
			put(cfg.Collector, "fee(uint256,uint256)", uintWord(1e17), uintWord(8), uintWord(feeTime))
			put(cfg.Burner, "want()", addrWord(cfg.Target))
			put(hook, "buffer_amount()", wad(20))
			duty := bytes.Join([][]byte{uintWord(32), addrWord(cfg.Allocator), uintWord(320), make([]byte, 7*32), uintWord(1), uintWord(4), common.RightPadBytes(pack("distribute_fees()"), 32)}, nil)
			put(hook, "hooks(uint256)", duty, uintWord(0))
			put(cfg.Allocator, "distributor_weight()", uintWord(9000))
			for h, n := range map[common.Address]int64{cfg.Collector: 100, cfg.Burner: 20, hook: 5, cfg.Allocator: 5} {
				put(cfg.Target, "balanceOf(address)", wad(n), addrWord(h))
			}
			put(lot, "balanceOf(address)", uintWord(0), addrWord(cfg.Collector))
			put(lot, "balanceOf(address)", uintWord(12), addrWord(cfg.Burner))
			put(cfg.Burner, "getAmountNeeded(address)", wad(8), addrWord(lot))
			server := rpc.NewServer()
			defer server.Stop()
			if err := server.RegisterName("eth", s); err != nil {
				t.Fatal(err)
			}
			httpServer := httptest.NewServer(server)
			defer httpServer.Close()
			t.Setenv("ETHEREUM_RPC_URL", httpServer.URL)
			catalogCalled := false
			row, err := readChain(context.Background(), cfg, options{pools: "auto"}, func() ([]interfaces.CurvePool, error) {
				catalogCalled = true
				return []interfaces.CurvePool{{BlockchainId: cfg.Name, Coins: []interfaces.Coin{{Address: lot.Hex(), Decimals: 0, UsdPrice: 2}, {Address: cfg.Target.Hex(), Decimals: 18, UsdPrice: 1}}}}, nil
			})
			if err != nil {
				t.Fatal(err)
			}
			if row.Status != "installed" || !row.RouteVerified || row.Phase != tt.phase || row.Block.Number != tt.head || row.Pools != nil || row.WeekFlows != nil || row.Pending.Partial {
				t.Fatalf("incorrect snapshot %+v", row)
			}
			p := row.Pending
			if p.EstimateCrvUSD == nil || *p.EstimateCrvUSD != tt.expected || p.VeEstimate == nil || *p.VeEstimate != tt.expected*0.9 || p.CashCrvUSD != "130" || p.BufferCrvUSD != "20" {
				t.Fatalf("incorrect amounts %+v", p)
			}
			if tt.phase == "forward" && (catalogCalled || row.InventoryScope != "crvUSD") {
				t.Fatalf("unnecessary inventory survey %+v", row)
			}
			if tt.phase == "exchange" && (!catalogCalled || row.InventoryScope != "crvUSDAndWeeklyLots" || p.CurrentLotQuote != "8" || p.RemainingMarketUSD != 20) {
				t.Fatalf("incorrect active lot %+v", row)
			}
		})
	}
}
