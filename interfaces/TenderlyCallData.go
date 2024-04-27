package interfaces

import "github.com/ethereum/go-ethereum/common"

type TenderlyCallData struct {
	From           common.Address `json:"from"`
	To             common.Address `json:"to"`
	Data           string         `json:"data"`
	SimulationType string         `json:"simulation_type"`
}
