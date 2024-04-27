package interfaces

type Claim struct {
	Timestamp uint64  `json:"timestamp"`
	Value     float64 `json:"value"`
}
