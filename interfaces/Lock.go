package interfaces

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Lock struct {
	Tx        common.Hash `json:"tx"`
	Timestamp uint64      `json:"timestamp"`
	Value     *big.Int    `json:"value"`
}
