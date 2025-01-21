package interfaces

import "github.com/ethereum/go-ethereum/common"

type ScrvusdDistribution struct {
	TxHash         common.Hash `json:"hash"`
	BlockTimestamp uint64      `json:"blockTimestamp"`
	Gain           float64     `json:"gain"`
	GainUSD        float64     `json:"gainUSD"`
}
