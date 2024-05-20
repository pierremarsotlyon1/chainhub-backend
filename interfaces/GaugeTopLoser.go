package interfaces

import "github.com/ethereum/go-ethereum/common"

type GaugeTopLoser struct {
	Gauge         common.Address `json:"gauge"`
	CurrentWeight float64        `json:"currentWeight"`
}
