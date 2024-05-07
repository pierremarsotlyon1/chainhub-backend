package interfaces

type VestingClaim struct {
	Timestamp uint64  `json:"ts"`
	AmountCRV float64 `json:"amountCRV"`
	AmountUSD float64 `json:"amountUSD"`
}
