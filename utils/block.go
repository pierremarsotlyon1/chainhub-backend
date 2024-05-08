package utils

import (
	"encoding/json"
	"io"
	"main/interfaces"
	"net/http"
	"strconv"
	"time"
)

func GetBlockNumberByTimestamp(chain string, timestamp uint64) uint64 {

	response, err := http.Get("https://coins.llama.fi/block/" + chain + "/" + strconv.FormatUint(timestamp, 10))

	// For limit rate
	time.Sleep(500 * time.Millisecond)

	if err != nil {
		return 0
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0
	}

	defilammaBlock := new(interfaces.DefilammaBlock)
	if err := json.Unmarshal(body, defilammaBlock); err != nil {
		return 0
	}

	return defilammaBlock.Height
}
