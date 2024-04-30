package interfaces

type StatsClaim struct {
	Claims                []Claim               `json:"claims"`
	TotalClaimed          float64               `json:"totalClaimed"`
	VeCRVTotalClaimed     float64               `json:"veCRVTotalClaimed"`
	VlCVXTotalClaimed     float64               `json:"vlCVXTotalClaimed"`
	VeStatsClaimBreakdown []StatsClaimBreakdown `json:"veBreakdown"`
}

type StatsClaimBreakdown struct {
	Name       string  `json:"name"`
	ClaimedUSD float64 `json:"claimedUsd"`
}
