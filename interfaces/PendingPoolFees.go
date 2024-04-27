package interfaces

type PendingPoolFees struct {
	LastBlock                 uint64             `json:"lastBlock"`
	LastDistributionTimestamp uint64             `json:"lastDistributionTime"`
	LastDistributionAmount    uint64             `json:"lastDistributionAmount"`
	Pending3CRVFees           map[string]float64 `json:"pending3CRVFees"`
}
