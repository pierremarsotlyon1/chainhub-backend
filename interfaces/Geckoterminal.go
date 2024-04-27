package interfaces

type Geckoterminal struct {
	Data struct {
		Attributes struct {
			PriceUSD string `json:"price_usd"`
		} `json:"attributes"`
	} `json:"data"`
}
