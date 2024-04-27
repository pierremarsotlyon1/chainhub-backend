package interfaces

type TenderlyForkResponse struct {
	SimulationFork struct {
		Id     string `json:"id"`
		RpcUrl string `json:"rpc_url"`
	} `json:"simulation_fork"`
}
