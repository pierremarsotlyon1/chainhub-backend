package interfaces

type CowswapQuote struct {
	Quote struct {
		BuyAmount string `json:"buyAmount"`
	} `json:"quote"`
}
