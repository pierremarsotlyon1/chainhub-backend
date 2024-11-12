package interfaces

type Coin struct {
	Address  string      `json:"address"`
	Symbol   string      `json:"symbol"`
	UsdPrice interface{} `json:"usdPrice"`
	Decimals interface{} `json:"decimals"`
}

type CurvePool struct {
	LpTokenAddress  string  `json:"lpTokenAddress"`
	GaugeAddress    string  `json:"gaugeAddress"`
	Address         string  `json:"address"`
	Coins           []Coin  `json:"coins"`
	UnderlyingCoins []Coin  `json:"underlyingCoins"`
	BlockchainId    string  `json:"blockchainId"`
	TotalSupply     string  `json:"totalSupply"`
	UsdTotal        float64 `json:"usdTotal"`
}

type GetAllCurvePoolsResponse struct {
	Data struct {
		PoolData []CurvePool `json:"poolData"`
	} `json:"data"`
}
