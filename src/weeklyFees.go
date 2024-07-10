package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/feeDistributor"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	weeklyFeesPath = "./data/weeklyFees/weekly-fees.json"
)

func WeeklyFees(client *ethclient.Client) {
	currentBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentPeriod := uint64(currentBlock.Time()/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC

	threeCrvPrice := utils.GetHistoricalPriceTokenPrice(utils.THREE_CRV, "ethereum", uint64(time.Now().Unix()))
	crvUSDPrice := 1.0 // assume peg to 1

	table := make([]interfaces.WeeklyFeesTable, 0)
	for _, feeDistributorAddress := range []common.Address{utils.FEE_DISTRIBUTOR_MAINNET, utils.FEE_DISTRIBUTOR_THREE_CRV_MAINNET} {
		feeDistributorContract, err := feeDistributor.NewFeeDistributor(feeDistributorAddress, client)
		if err != nil {
			fmt.Println(err)
			return
		}

		startTime, err := feeDistributorContract.StartTime(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		period := uint64(uint64(startTime.Int64())/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC
		run := true

		for run {
			tokenPerWeek, err := feeDistributorContract.TokensPerWeek(nil, big.NewInt(int64(period)))
			if err == nil {
				fees := utils.Quo(tokenPerWeek, 18)
				if fees > 0 {
					price := 0.0
					symbol := ""
					if strings.EqualFold(feeDistributorAddress.Hex(), utils.FEE_DISTRIBUTOR_THREE_CRV_MAINNET.Hex()) {
						price = threeCrvPrice
						symbol = "3CRV"
					} else {
						price = crvUSDPrice
						symbol = "crvUSD"
					}

					table = append(table, interfaces.WeeklyFeesTable{
						Ts:      period,
						Fees:    fees,
						FeesUSD: fees * price,
						Symbol:  symbol,
					})
				}
			}

			period += utils.WEEK_TO_SEC

			if period > currentPeriod {
				run = false
				break
			}
		}
	}

	writeWeeklyFees(table)
}

func readWeeklyFees() []interfaces.WeeklyFeesTable {

	if !utils.FileExists(weeklyFeesPath) {
		return make([]interfaces.WeeklyFeesTable, 0)
	}

	file, err := os.ReadFile(weeklyFeesPath)
	if err != nil {
		log.Fatal(err)
	}

	tables := make([]interfaces.WeeklyFeesTable, 0)
	if err := json.Unmarshal([]byte(file), &tables); err != nil {
		log.Fatal(err)
	}

	return tables
}

func writeWeeklyFees(tables []interfaces.WeeklyFeesTable) {

	// Sum fees per period
	newTable := make([]interfaces.WeeklyFeesTable, 0)

	for _, t := range tables {

		found := false
		for i, t2 := range newTable {
			if t2.Ts == t.Ts {
				newTable[i].Fees += t.Fees
				newTable[i].FeesUSD += t.FeesUSD
				found = true
				break
			}
		}

		if !found {
			newTable = append(newTable, t)
		}
	}

	sort.Slice(newTable, func(i, j int) bool { return newTable[i].Ts > newTable[j].Ts })

	file, err := json.Marshal(newTable)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(weeklyFeesPath, file, 0644); err != nil {
		log.Fatal(err)
	}
}
