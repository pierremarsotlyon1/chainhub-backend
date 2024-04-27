package interfaces

type CrvUSDHistoricalData struct {
	TotalSupply float64 `json:"totalSupply"`
	Timestamp   uint64  `json:"timestamp"`
	Tvl         float64 `json:"tvl"`
}

type CrvUSDHistorical struct {
	CurrentCrvUSDPrice  float64                `json:"currentCrvUsdPrice"`
	DailyHistoricalData []CrvUSDHistoricalData `json:"dailyHistoricalData"`
}
