package interfaces

type AlchemyTransfer struct {
	BlockNum    string  `json:"blockNum"`
	Hash        string  `json:"hash"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	Value       float64 `json:"value"`
	RawContract struct {
		Address string `json:"address"`
		Decimal string `json:"decimal"`
		Value   string `json:"value"`
	} `json:"rawContract"`
}

type AlchemyTransfers struct {
	Result struct {
		Transfers []AlchemyTransfer `json:"transfers"`
		PageKey   string            `json:"pageKey"`
	} `json:"result"`
}
