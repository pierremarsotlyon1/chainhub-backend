package interfaces

type VeCrvFees struct {
	Components map[string]float64 `json:"components"` // USD attributable to veCRV
	Partial    bool               `json:"partial"`
}

type PendingPoolFees struct {
	VeCrvFees                 VeCrvFees          `json:"veCrvFees"`
	LastBlock                 uint64             `json:"lastBlock"`
	LastDistributionTimestamp uint64             `json:"lastDistributionTime"`
	LastDistributionAmount    uint64             `json:"lastDistributionAmount"`
	Pending3CRVFees           map[string]float64 `json:"pending3CRVFees"`
}
