package interfaces

type InflationPerMonth struct {
	Timestamp                 int64   `json:"ts"`
	BountiesIncomeAmountUSD   float64 `json:"bountiesIncomeAmountUSD"`
	WeeklyFeesIncomeAmountUSD float64 `json:"weeklyFeesIncomeAmountUSD"`
	TotalIncomeAmountUSD      float64 `json:"totalIncomeAmountUSD"`
	InflationUSD              float64 `json:"inflationUSD"`
	InflationCRV              float64 `json:"inflationCRV"`
	VestingUSD                float64 `json:"vestingUSD"`
	VestingCRV                float64 `json:"vestingCRV"`
}
