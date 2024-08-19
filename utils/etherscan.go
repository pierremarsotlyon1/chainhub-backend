package utils

import (
	"io"
	"net/http"
	"strings"
)

func ContainsWithdrawAdminFees(contractAddress string) (bool, error) {

	response, err := http.Get("https://api.etherscan.io/api?module=contract&action=getabi&address=" + contractAddress + "&apikey=" + GoDotEnvVariable("ETHERSCAN_API_KEY"))
	if err != nil {
		return false, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(body), "withdraw_admin_fees"), nil
}
