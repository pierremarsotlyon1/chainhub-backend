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
		return "https://opt-mainnet.g.alchemy.com/v2/" + GoDotEnvVariable("ALCHEMY_APIKEY_2")
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
	default:
		return ""
	}
}
