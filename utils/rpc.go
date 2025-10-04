package utils

func GetPublicRpcUrl(chainName string) string {
	switch chainName {
	case "polygon":
		return "https://rpc.crvhub.com/main/evm/137"
	case "arbitrum":
		return "https://rpc.crvhub.com/main/evm/42161"
	case "base":
		return "https://rpc.crvhub.com/main/evm/8453"
	case "celo":
		return "https://rpc.crvhub.com/main/evm/42220"
	case "optimism":
		return "https://rpc.crvhub.com/main/evm/10"
	case "xdai":
		return "https://rpc.crvhub.com/main/evm/100"
	case "avalanche":
		return "https://rpc.crvhub.com/main/evm/43114"
	case "fantom":
		return "https://rpc.crvhub.com/main/evm/250"
	case "kava":
		return "https://rpc.crvhub.com/main/evm/2222"
	case "moonbeam":
		return "https://rpc.crvhub.com/main/evm/1284"
	case "fraxtal":
		return "https://rpc.crvhub.com/main/evm/252"
	case "bsc":
		return "https://rpc.crvhub.com/main/evm/56"
	case "x-layer":
		return "https://rpc.crvhub.com/main/evm/196"
	case "sonic":
		return "https://rpc.crvhub.com/main/evm/146"
	case "taiko":
		return "https://rpc.crvhub.com/main/evm/167000"
	case "mainnet":
		return "https://rpc.crvhub.com/main/evm/1"
	default:
		return ""
	}
}
