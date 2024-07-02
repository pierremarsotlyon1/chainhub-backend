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
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	weeklyFeesConfig       = "./data/configs/weekly-fees-config.json"
	transferWeeklyFeesPath = "./data/weeklyFees/transfers.json"
	weeklyFeesPath         = "./data/weeklyFees/weekly-fees.json"
)

func WeeklyFees(client *ethclient.Client, currentBlock uint64) {

	config := utils.ReadConfig(weeklyFeesConfig)

	transfers := readWeeklyFeesTransfers()
	transfers = append(transfers, fetchWeeklyFees(client, currentBlock, config)...)

	utils.WriteConfig(config, currentBlock, weeklyFeesConfig)
	writeWeeklyFeesTransfers(transfers)
	computeTransfer(transfers)
}

func fetchWeeklyFees(client *ethclient.Client, currentBlock uint64, config interfaces.Config) []interfaces.WeeklyFee {

	from := config.LastBlock
	if from == 0 {
		from = 11278886
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{utils.FEE_DISTRIBUTOR_MAINNET},
		Topics:    [][]common.Hash{{common.HexToHash("0xce749457b74e10f393f2c6b1ce4261b78791376db5a3f501477a809f03f500d6")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	transfers := make([]interfaces.WeeklyFee, 0)

	for _, vLog := range logs {

		feeDistributorContract, err := feeDistributor.NewFeeDistributor(vLog.Address, client)
		if err != nil {
			panic(err)
		}

		event, err := feeDistributorContract.ParseCheckpointToken(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		header, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			continue
		}

		if event.Tokens.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		transfers = append(transfers, interfaces.WeeklyFee{
			Ts:     header.Time(),
			Fees:   utils.Quo(event.Tokens, 18),
			TxHash: vLog.TxHash,
		})
	}

	return transfers
}

func computeTransfer(transfers []interfaces.WeeklyFee) {
	table := make([]interfaces.WeeklyFeesTable, 0)

	for _, transfer := range transfers {
		currentPeriod := uint64(transfer.Ts/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC

		found := false
		for i := 0; i < len(table); i++ {
			if table[i].Ts == currentPeriod {
				table[i].Fees += transfer.Fees
				found = true
				break
			}
		}

		if !found {
			table = append(table, interfaces.WeeklyFeesTable{
				Ts:   currentPeriod,
				Fees: transfer.Fees,
			})
		}
	}

	threeCrvPrice := utils.GetHistoricalPriceTokenPrice(utils.THREE_CRV, "ethereum", uint64(time.Now().Unix()))
	crvUSDPrice := utils.GetHistoricalPriceTokenPrice(utils.CRVUSD_ADDRESS, "ethereum", uint64(time.Now().Unix()))

	for i := 0; i < len(table); i++ {
		if table[i].Ts >= 1718841600 {
			// crvUSD
			table[i].FeesUSD = table[i].Fees * crvUSDPrice
			table[i].Symbol = "crvUSD"
		} else {
			// 3CRV
			table[i].FeesUSD = table[i].Fees * threeCrvPrice
			table[i].Symbol = "3CRV"
		}
	}

	writeWeeklyFees(table)
}

func readWeeklyFeesTransfers() []interfaces.WeeklyFee {

	if !utils.FileExists(transferWeeklyFeesPath) {
		return make([]interfaces.WeeklyFee, 0)
	}

	file, err := os.ReadFile(transferWeeklyFeesPath)
	if err != nil {
		log.Fatal(err)
	}

	transfers := make([]interfaces.WeeklyFee, 0)
	if err := json.Unmarshal([]byte(file), &transfers); err != nil {
		log.Fatal(err)
	}

	return transfers
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

func writeWeeklyFeesTransfers(transfers []interfaces.WeeklyFee) {
	file, err := json.Marshal(transfers)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(transferWeeklyFeesPath, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func writeWeeklyFees(tables []interfaces.WeeklyFeesTable) {
	file, err := json.Marshal(tables)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(weeklyFeesPath, file, 0644); err != nil {
		log.Fatal(err)
	}
}
