package interfaces

type TenderlyForkTxListResponse struct {
	ForkTransactions []struct {
		Id        string `json:"id"`
		ProjectId string `json:"project_id"`
		ForkId    string `json:"fork_id"`
	} `json:"fork_transactions"`
}
