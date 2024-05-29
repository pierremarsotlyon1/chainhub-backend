package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
)

func GetIpfs(ipfsId string) string {
	response, err := http.Get("https://gateway.pinata.cloud/ipfs/" + ipfsId)
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}

	ipfs := new(interfaces.Ipfs)
	if err := json.Unmarshal(body, ipfs); err != nil {
		return ""
	}

	return ipfs.Text
}
