package interfaces

type Stats struct {
	TotalClaimed         float64    `json:"totalClaimed"`
	VeCRVTotalClaimed    float64    `json:"veCRVTotalClaimed"`
	VlCVXTotalClaimed    float64    `json:"vlCVXTotalClaimed"`
	ClaimsLast7Days      StatsClaim `json:"claimsLast7Days"`
	ClaimsLast30Days     StatsClaim `json:"claimsLast30Days"`
	ClaimsLast180Days    StatsClaim `json:"claimsLast180Days"`
	ClaimsLast365Days    StatsClaim `json:"claimsLast365Days"`
	ClaimsSinceInception StatsClaim `json:"claimsSinceInception"`
}
