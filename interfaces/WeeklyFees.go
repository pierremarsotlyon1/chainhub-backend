package interfaces

type WeeklyFees struct {
	Data struct {
		WeeklyFeesTable []WeeklyFeesTable `json:"weeklyFeesTable"`
	} `json:"data"`
}

type WeeklyFeesTable struct {
	Date string  `json:"date"`
	Ts   uint64  `json:"ts"`
	Fees float64 `json:"rawFees"`
}
