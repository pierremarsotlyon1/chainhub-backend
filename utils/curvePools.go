package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
)

func GetAllCurvePools() ([]interfaces.CurvePool, error) {
	response, err := http.Get("https://api.curve.finance/v1/getPools/all")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	allCurvePools := new(interfaces.GetAllCurvePoolsResponse)
	if err := json.Unmarshal(body, allCurvePools); err != nil {
		return nil, err
	}

	return allCurvePools.Data.PoolData, nil
}
