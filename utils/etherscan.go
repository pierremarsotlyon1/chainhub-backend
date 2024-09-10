package utils

import (
	"io"
	"net/http"
	"strings"
)

func GetEtherscanAbi(contractAddress string) (string, error) {
	response, err := http.Get("https://api.etherscan.io/api?module=contract&action=getabi&address=" + contractAddress + "&apikey=" + GoDotEnvVariable("ETHERSCAN_API_KEY"))
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ContainsWithdrawAdminFees(abi string) (bool, error) {
	return strings.Contains(abi, "withdraw_admin_fees"), nil
}

func ContainsClaimAdminFees(abi string) (bool, error) {
	return strings.Contains(abi, "claim_admin_fees"), nil
}
