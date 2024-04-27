package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BountyClaimed struct {
	Contract      common.Address `json:"contract"`
	TokenReward   common.Address `json:"tokenReward"`
	TokenDecimals uint8          `json:"tokenDecimals"`
	Amount        *big.Int       `json:"amount"`
	Timestamp     uint64         `json:"timestamp"`
	BlockNumber   uint64         `json:"blocknumber"`
	Tx            common.Hash    `json:"tx"`
	Price         float64        `json:"historicalPrice"`
	Comment       string         `json:"comment"`
}
