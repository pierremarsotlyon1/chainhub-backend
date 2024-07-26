package interfaces

type BountyClaimed struct {
	Timestamp    uint64  `json:"t"`
	Comment      string  `json:"c"`
	AmountUSD    float64 `json:"a"`
	ProtocolName string  `json:"p"`
}
