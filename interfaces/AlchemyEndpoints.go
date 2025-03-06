package interfaces

type AlchemyEndpoints struct {
	Result struct {
		Data []AlchemyEndpoint `json:"data"`
	} `json:"result"`
}

type AlchemyEndpoint struct {
	NetworkChainId int64  `json:"networkChainId"`
	KebabCaseId    string `json:"kebabCaseId"`
}
