package utils

func GetPublicRpcUrl(chainName string) string {
	switch chainName {
	case "polygon":
		return "https://polygon-rpc.com"
	case "arbitrum":
		return "https://arb1.arbitrum.io/rpc"
	case "base":
		return "https://mainnet.base.org"
	case "celo":
		return "https://celo-mainnet.infura.io/v3/" + GoDotEnvVariable("INFURA_APIKEY")
	case "optimism":
		return "https://lb.drpc.org/optimism/" + GoDotEnvVariable("DRPC_API_KEY")
	case "xdai":
		return "https://gnosis-mainnet.public.blastapi.io"
	case "avalanche":
		return "https://avalanche-mainnet.infura.io/v3/" + GoDotEnvVariable("INFURA_APIKEY")
	case "fantom":
		return "https://rpc.fantom.network"
	case "kava":
		return "https://evm.kava.io"
	case "moonbeam":
		return "https://rpc.ankr.com/moonbeam"
	case "fraxtal":
		return "https://rpc.frax.com"
	case "bsc":
		return "https://rpc.ankr.com/bsc"
	case "x-layer":
		return "https://rpc.xlayer.tech"
	case "sonic":
		return "https://lb.drpc.org/sonic/" + GoDotEnvVariable("DRPC_API_KEY")
	case "taiko":
		return "https://lb.drpc.org/taiko/" + GoDotEnvVariable("DRPC_API_KEY")
	default:
		return ""
	}
}
