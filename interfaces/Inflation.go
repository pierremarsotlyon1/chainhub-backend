package interfaces

type InflationData struct {
	Inflation []Inflation `json:"inflations"`
}

type Inflation struct {
	Timestamp int64   `json:"ts"`
	AmountCRV float64 `json:"amountCRV"`
	AmoutUSD  float64 `json:"amountUSD"`
}
