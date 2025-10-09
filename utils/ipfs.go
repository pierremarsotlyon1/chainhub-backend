package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
)

var providers = []string{"https://4everland.io/ipfs/", "https://gateway.pinata.cloud/ipfs/", "https://ipfs.io/ipfs/"}

func GetIpfs(ipfsId string) string {
	for _, provider := range providers {
		description := fetchIpfs(ipfsId, provider)
		if len(description) > 0 {
			return description
		}
	}

	return ""
}

func fetchIpfs(ipfsId string, provider string) string {
	response, err := http.Get(provider + ipfsId)
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

	if len(ipfs.Text) > 0 {
		return ipfs.Text
	}

	if len(ipfs.Description) > 0 {
		return ipfs.Description
	}

	return ""
}
