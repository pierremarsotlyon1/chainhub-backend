package dutchfees

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	cfg := chainConfig{Burner: addr(1), Allocator: addr(2), Bridger: addr(3), Treasury: addr(4)}
	target, hook, token := addr(5), addr(6), addr(7)
	logs := []types.Log{
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
	if staged[token].Int64() != 22 || f.NativeTakenCrvUSD != want(15) || f.CowTradeCrvUSD != want(18) || f.RealizedCrvUSD != want(33) || f.DistributedGrossCrvUSD != want(45) || f.DistributedVeCRVCrvUSD != want(35) || f.BridgeInitiatedCrvUSD != want(70) {
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
