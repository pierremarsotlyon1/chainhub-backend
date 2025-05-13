package src

import (
	"fmt"
	"io"
	"main/utils"
	"net/http"
)

var endpoints = []string{
	"https://api.curve.finance/v1/getPools/all",
	"https://api-core.curve.finance/v1/getPlatforms",
}

var endpointsBucket = []string{
	"getPoolsAll",
	"getPlatforms",
}

func CurveApiCache() {
	for i := range endpoints {

		response, err := http.Get(endpoints[i])

		if err != nil {
			continue
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			continue
		}

		utils.WriteBucketFile(fmt.Sprintf("data/curve_api/%s.json", endpointsBucket[i]), string(body))
	}
}
