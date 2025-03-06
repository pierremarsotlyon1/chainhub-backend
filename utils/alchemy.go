package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
)

func FetchAllAlchemyEndpoints() []interfaces.AlchemyEndpoint {

	response, err := http.Get("https://app-api.alchemy.com/trpc/config.getNetworkConfig")

	if err != nil {
		return []interfaces.AlchemyEndpoint{}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []interfaces.AlchemyEndpoint{}
	}

	alchemyEndpoints := new(interfaces.AlchemyEndpoints)
	if err := json.Unmarshal(body, alchemyEndpoints); err != nil {
		return []interfaces.AlchemyEndpoint{}
	}

	return alchemyEndpoints.Result.Data
}
