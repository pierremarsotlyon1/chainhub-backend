package interfaces

type DecodedAction struct {
	Index      int             `json:"index"`
	Target     string          `json:"target"`
	Function   string          `json:"function"`
	Args       map[string]any  `json:"args"`
	SubActions []DecodedAction `json:"subActions,omitempty"`
}
