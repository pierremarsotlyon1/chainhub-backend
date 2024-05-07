package src

import (
	"encoding/json"
	"log"
	"main/contracts/crv"
	"main/interfaces"
	"main/utils"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	INFLATION_PATH = "./data/inflation/data.json"
)

func Inflation(client *ethclient.Client, currentBlockf uint64, currentBlockTImestamp uint64) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	inflationData := readInflation()

	// Check if we already got the inflation
	if len(inflationData.Inflation) > 0 && inflationData.Inflation[len(inflationData.Inflation)-1].Timestamp == startOfDay {
		return
	}

	crvContract, err := crv.NewCrv(utils.CRV, client)
	if err != nil {
		log.Fatal(err)
	}

	rateBN, err := crvContract.Rate(nil)
	if err != nil {
		log.Fatal(err)
	}

	rate := utils.Quo(rateBN, 18)
	amountCRV := rate * float64(utils.DAY_TO_SEC)
	amountUSD := amountCRV * utils.GetHistoricalPriceTokenPrice(utils.CRV, "ethereum", uint64(startOfDay))

	inflationData.Inflation = append(inflationData.Inflation, interfaces.Inflation{
		Timestamp: startOfDay,
		AmountCRV: amountCRV,
		AmoutUSD:  amountUSD,
	})

	writeInflation(inflationData)
}

func InflationHistorical(client *ethclient.Client, currentBlockf uint64, currentBlockTImestamp uint64) {
	startOfDay := uint64(1597269600)
	now := uint64(time.Now().Unix())

	inflationData := readInflation()

	for {

		rate := 0.0
		if startOfDay < 1628893332 {
			rate = 8.71433545789
		} else if startOfDay < 1660429052 {
			rate = 7.32785344786
		} else if startOfDay < 1691965151 {
			rate = 6.16196569581
		} else {
			rate = 5.18157486452
		}

		amountCRV := rate * float64(utils.DAY_TO_SEC)
		amountUSD := amountCRV * utils.GetHistoricalPriceTokenPrice(utils.CRV, "ethereum", uint64(startOfDay))

		inflationData.Inflation = append(inflationData.Inflation, interfaces.Inflation{
			Timestamp: int64(startOfDay),
			AmountCRV: amountCRV,
			AmoutUSD:  amountUSD,
		})

		startOfDay += utils.DAY_TO_SEC
		if startOfDay > now {
			break
		}
	}

	writeInflation(inflationData)
}

func readInflation() interfaces.InflationData {

	if !utils.FileExists(INFLATION_PATH) {
		return interfaces.InflationData{
			Inflation: make([]interfaces.Inflation, 0),
		}
	}

	file, err := os.ReadFile(INFLATION_PATH)
	if err != nil {
		log.Fatal(err)
	}

	inflation := new(interfaces.InflationData)
	if err := json.Unmarshal([]byte(file), &inflation); err != nil {
		log.Fatal(err)
	}

	return *inflation
}

func writeInflation(inflation interfaces.InflationData) {
	file, err := json.Marshal(inflation)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(INFLATION_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}
