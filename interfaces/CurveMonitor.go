package interfaces

type CurveMonitor struct {
	Proposals []struct {
		VoteId       int64  `json:"voteId"`
		IpfsMetadata string `json:"ipfsMetadata"`
		Metadata     string `json:"metadata"`
	} `json:"proposals`
}
