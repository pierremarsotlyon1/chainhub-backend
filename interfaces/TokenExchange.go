package interfaces

import "github.com/ethereum/go-ethereum/common"

type TokenExchange struct {
	PegKeeperAddress common.Address `json:"pegkeeperAddress"`
	Amount           float64        `json:"amount"`
	LpPrice          float64        `json:"lpPrice"`
	TxHash           common.Hash    `json:"txHash"`
	BlockNumber      uint64         `json:"blockNumber"`
	BlockTimestamp   uint64         `json:"blockTimestamp"`
}
