package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
)

type GaugeResp struct {
	Data map[string]interfaces.Gauge `json:"data"`
}

func GetAllGauges() []interfaces.Gauge {

	gauges := make([]interfaces.Gauge, 0)
	response, err := http.Get("https://api.curve.finance/api/getAllGauges")
	if err != nil {
		return gauges
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return gauges
	}

	gaugeResp := new(GaugeResp)
	if err := json.Unmarshal(body, gaugeResp); err != nil {
		return gauges
	}

	for _, gauge := range gaugeResp.Data {
		gauges = append(gauges, gauge)
	}

	return gauges
}
