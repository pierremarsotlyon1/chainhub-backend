package interfaces

import "github.com/ethereum/go-ethereum/common"

type LlamalendLiquidate struct {
	Timestamp uint64         `json:"ts"`
	User      common.Address `json:"user"`
}
