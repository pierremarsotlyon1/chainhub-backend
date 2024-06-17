package interfaces

type StatsLock struct {
	Amount    float64 `json:"amount"`
	Timestamp uint64  `json:"timestamp"`
	StakeDAO  float64 `json:"stakedao"`
	Yearn     float64 `json:"yearn"`
	Convex    float64 `json:"convex"`
}
