package interfaces

type TenderlyBatchSimulation struct {
	Id      int64         `json:"id"`
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
