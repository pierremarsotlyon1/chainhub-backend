package interfaces

import "github.com/ethereum/go-ethereum/common"

type PegKeeper struct {
	Address       common.Address `json:"address"`
	PoolAddress   common.Address `json:"poolAddress"`
	Name          string         `json:"name"`
	Tvl           float64        `json:"tvl"`
	Debt          float64        `json:"debt"`
	FeesCollected float64        `json:"feesCollected"`
	FeesPending   float64        `json:"feesPending"`
}
