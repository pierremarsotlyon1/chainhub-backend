package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"math/big"
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

func GetIpfsFromCurveMonitor(ipfsId string, voteId *big.Int) string {
	response, err := http.Get("https://api-py.llama.airforce/curve/v1/dao/proposals")
	if err != nil {
		return ""
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}

	curveMonitor := new(interfaces.CurveMonitor)
	if err := json.Unmarshal(body, curveMonitor); err != nil {
		return ""
	}

	for _, proposal := range curveMonitor.Proposals {
		if proposal.VoteId == voteId.Int64() {
			return proposal.Metadata
		}
	}
	return ""
}
