package dutchfees

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestWeeklyFlowAccounting(t *testing.T) {
	addr := func(n int64) common.Address { return common.BigToAddress(big.NewInt(n)) }
	topic := func(a common.Address) common.Hash { return common.BytesToHash(a.Bytes()) }
	data := func(values ...int64) []byte {
		var out []byte
		for _, v := range values {
			out = append(out, common.LeftPadBytes(big.NewInt(v).Bytes(), 32)...)
		}
		return out
	}
	cfg := chainConfig{LegacyBurner: addr(10), Burner: addr(1), Allocator: addr(2), Bridger: addr(3), Treasury: addr(4)}
	target, hook, token := addr(5), addr(6), addr(7)
	logs := []types.Log{
		{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, topic(cfg.LegacyBurner)}, Data: data(7, 5, 99, 11, 2, 192, 0)},
		{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, topic(cfg.LegacyBurner)}, Data: data(7, 8, 99, 27, 2, 192, 0)},
		{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, topic(addr(11))}, Data: data(7, 5, 99, 44, 2, 192, 0)},
		{Address: cfg.Burner, Topics: []common.Hash{lotStagedEvent, topic(token)}, Data: data(20, 1, 1, 1, 1)},
		{Address: cfg.Burner, Topics: []common.Hash{nativeTakenEvent, topic(token), topic(hook)}, Data: data(9, 100, 15, 3)},
		{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, topic(cfg.Burner)}, Data: data(7, 5, 99, 18, 2, 192, 0)},
		{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, topic(cfg.Burner)}, Data: data(7, 8, 99, 55, 2, 192, 0)},
		{Address: cfg.Allocator, Topics: []common.Hash{feeDistributionEvent}, Data: data(45, 35)},
		{Address: cfg.Bridger, Topics: []common.Hash{bridgeInitiatedEvent, topic(target), topic(hook), topic(cfg.Treasury)}, Data: data(70)},
		{Address: cfg.Bridger, Topics: []common.Hash{bridgeInitiatedEvent, topic(target), topic(addr(9)), topic(cfg.Treasury)}, Data: data(99)},
		{Address: cfg.Burner, Topics: []common.Hash{lotStagedEvent, topic(token)}, Data: data(22, 1, 1, 1, 1)},
	}
	for i := range logs {
		logs[i].Index = uint(i)
	}
	logs = append(logs, logs[1])
	f, staged, err := summarizeFlowLogs(cfg, target, hook, logs)
	if err != nil {
		t.Fatal(err)
	}
	want := func(n int64) string { return amountText(big.NewInt(n), 18) }
	if staged[token].Int64() != 22 || f.NativeTakenCrvUSD != want(15) || f.CowTradeCrvUSD != want(29) || f.RealizedCrvUSD != want(44) || f.DistributedGrossCrvUSD != want(45) || f.DistributedVeCRVCrvUSD != want(35) || f.BridgeInitiatedCrvUSD != want(70) {
		t.Fatalf("unexpected flows %+v, staged %v", f, staged)
	}
}
func TestWeeklyFlowRejectsReorg(t *testing.T) {
	log := types.Log{Topics: []common.Hash{lotStagedEvent}, Removed: true}
	if _, _, err := summarizeFlowLogs(chainConfig{}, common.Address{}, common.Address{}, []types.Log{log}); err == nil {
		t.Fatal("accepted removed log")
	}
}

type gatewayRangeError struct{}

func (gatewayRangeError) Error() string { return "getLogs request exceeded max allowed range" }
func (gatewayRangeError) ErrorData() interface{} {
	return map[string]interface{}{"details": map[string]interface{}{"maxAllowedRange": 30000}}
}
func TestFlowProviderRange(t *testing.T) {
	err := gatewayRangeError{}
	if !flowRangeLimit(err) || flowProviderLimit(err) != 30000 {
		t.Fatal("did not use advertised gateway page size")
	}
	if !flowRangeLimit(errors.New("413 Request Entity Too Large")) || flowRangeLimit(errors.New("unauthorized")) {
		t.Fatal("incorrect fallback pagination")
	}
}

type weeklyFlowRPC struct {
	flowHeaderStub
	cfg chainConfig
}

func (s weeklyFlowRPC) GetLogs(q struct {
	FromBlock, ToBlock string
	Address            []common.Address
	Topics             [][]common.Hash
}) ([]types.Log, error) {
	if !slices.Contains(q.Address, cowSettlement) {
		return []types.Log{}, nil // No Dutch lot staged.
	}
	owners := []common.Hash{common.BytesToHash(s.cfg.LegacyBurner.Bytes()), common.BytesToHash(s.cfg.Burner.Bytes())}
	if q.FromBlock != "0x1f4" || q.ToBlock != "0x257" || len(q.Topics) != 2 || !slices.Equal(q.Topics[0], []common.Hash{cowTradeEvent}) || !slices.Equal(q.Topics[1], owners) {
		return nil, fmt.Errorf("unexpected settlement filter %+v", q)
	}
	data := make([]byte, 224)
	copy(data[32:64], addrWord(s.cfg.Target))
	copy(data[96:128], uintWord(19))
	return []types.Log{{Address: cowSettlement, Topics: []common.Hash{cowTradeEvent, owners[0]}, Data: data, BlockNumber: 500}}, nil
}

func TestWeeklyFlowsWithoutDutchLot(t *testing.T) {
	cfg := networks[0]
	server := rpc.NewServer()
	defer server.Stop()
	if err := server.RegisterName("eth", weeklyFlowRPC{cfg: cfg}); err != nil {
		t.Fatal(err)
	}
	r := &reader{client: ethclient.NewClient(rpc.DialInProc(server))}
	defer r.client.Close()
	head := snapshot{Number: 980, Timestamp: 298}
	p := protocol{Target: cfg.Target, CollectEnd: 250, ExchangeEnd: 260}
	f, staged, err := weeklyFlows(context.Background(), r, cfg, 0, head, p)
	if err != nil || len(staged) != 0 || f.RealizedCrvUSD != amountText(big.NewInt(19), 18) {
		t.Fatalf("weekly flows %+v, staged %v, err %v", f, staged, err)
	}
}
