package interfaces

type Coin struct {
	Address  string      `json:"address"`
	Symbol   string      `json:"symbol"`
	UsdPrice interface{} `json:"usdPrice"`
}

type CurvePool struct {
	LpTokenAddress  string `json:"lpTokenAddress"`
	GaugeAddress    string `json:"gaugeAddress"`
	Address         string `json:"address"`
	Coins           []Coin `json:"coins"`
	UnderlyingCoins []Coin `json:"underlyingCoins"`
	BlockchainId    string `json:"blockchainId"`
}

type GetAllCurvePoolsResponse struct {
	Data struct {
		PoolData []CurvePool `json:"poolData"`
	} `json:"data"`
}
