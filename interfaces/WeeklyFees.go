package interfaces

import "github.com/ethereum/go-ethereum/common"

type WeeklyFees struct {
	Data struct {
		WeeklyFeesTable []WeeklyFeesTable `json:"weeklyFeesTable"`
	} `json:"data"`
}

type WeeklyFeesTable struct {
	Ts      uint64  `json:"ts"`
	Fees    float64 `json:"rawFees"`
	FeesUSD float64 `json:"feesUSD"`
	Symbol  string  `json:"symbol"`
}

type WeeklyFee struct {
	Ts     uint64      `json:"ts"`
	Fees   float64     `json:"rawFees"`
	TxHash common.Hash `json:"tx"`
}
