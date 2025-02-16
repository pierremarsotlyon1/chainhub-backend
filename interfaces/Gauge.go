package interfaces

type Gauge struct {
	Pool                string  `json:"swap"`
	Gauge               string  `json:"gauge"`
	LendingVaultAddress string  `json:"lendingVaultAddress"`
	LP                  string  `json:"swap_token"`
	Name                string  `json:"name"`
	IsKilled            bool    `json:"is_killed"`
	HasNoCrv            bool    `json:"hasNoCrv"`
	CurrentMinApr       float64 `json:"currentMinApr"`
	CurrentMaxApr       float64 `json:"currentMaxApr"`
	NextMinApr          float64 `json:"nextMinApr"`
	NextMaxApr          float64 `json:"nextMaxApr"`
	SideChain           bool    `json:"side_chain"`
	GaugeController     struct {
		GaugeRelativeWeight       string `json:"gauge_relative_weight"`
		GaugeFutureRelativeWeight string `json:"gauge_future_relative_weight"`
	} `json:"gauge_controller"`
	IsPool            bool                `json:"isPool"`
	GaugeCrvApy       []float64           `json:"gaugeCrvApy"`
	GaugeFutureCrvApy []float64           `json:"gaugeFutureCrvApy"`
	PoolUrls          map[string][]string `json:"poolUrls"`
	LpTokenPrice      float64             `json:"lpTokenPrice"`
	PendingFees       float64             `json:"pendingFees"`
	AmountStakedUSD   float64             `json:"amountStakedUSD"`

	WorkingSupply       float64     `json:"workingSupply"`
	VirtualPrice        interface{} `json:"virtualPrice"`
	CanCalculateInTheUI bool        `json:"canCalculateInTheUI"`
}
