package interfaces

type DefilammaPrice struct {
	Coins map[string]struct {
		Price float64 `json:"price"`
	} `json:"coins"`
}
