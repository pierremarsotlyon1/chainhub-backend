package interfaces

import "github.com/ethereum/go-ethereum/common"

type PoolTokenExchange struct {
	TxHash             common.Hash    `json:"hash"`
	User               common.Address `json:"user"`
	TokenSoldAddress   common.Address `json:"tokenSoldAddress"`
	TokenSoldAmount    float64        `json:"tokenSoldAmount"`
	SymbolSold         string         `json:"symbolSold"`
	TokenBoughtAddress common.Address `json:"tokenBoughtAddress"`
	TokenBoughtAmount  float64        `json:"tokenBoughtAmount"`
	SymbolBought       string         `json:"symbolBought"`
	BlockTimestamp     uint64         `json:"blockTimestamp"`
	PoolAddress        common.Address `json:"poolAddress"`
}
