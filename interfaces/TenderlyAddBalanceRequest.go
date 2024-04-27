package interfaces

type TenderlyAddBalanceRequest struct {
	Id      int64         `json:"id"`
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
