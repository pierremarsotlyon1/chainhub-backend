package interfaces

type TenderlyBatchSimulationResponse struct {
	ID      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Result  []Result `json:"result"`
}

type Result struct {
	Trace []Trace `json:"trace"`
}

type Trace struct {
	DecodedOutput []DecodedOutput `json:"decodedOutput,omitempty"`
}

type DecodedOutput struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
	Name  string      `json:"name"`
}
