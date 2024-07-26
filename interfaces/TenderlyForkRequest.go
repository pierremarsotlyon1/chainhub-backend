package interfaces

type TenderlyForkRequest struct {
	NetworkId   string `json:"network_id"`
	BlockNumber int64  `json:"block_number"`
	Shared      bool   `json:"shared"`
}
