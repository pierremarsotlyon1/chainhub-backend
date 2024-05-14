package interfaces

import "github.com/ethereum/go-ethereum/common"

type Peg struct {
	Timestamp   uint64         `json:"timestamp"`
	PoolAddress common.Address `json:"poolAddress"`
	Peg         float64        `json:"peg"`
}
