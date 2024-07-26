package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/crv"
	"main/contracts/vesting"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	INFLATION_PATH              = "./data/inflation/data.json"
	INFLATION_PER_MONTH_PATH    = "./data/inflation/inflation-per-month.json"
	VESTING_CLAIMED_PATH_TMP    = "./data/inflation/vesting-claimed-tmp.json"
	VESTING_CLAIMED_PATH        = "./data/inflation/vesting-claimed.json"
	VESTING_CLAIMED_CONFIG_PATH = "./data/configs/vesting-claimed.json"
)

func Inflation(client *ethclient.Client, currentBlock uint64, currentBlockTImestamp uint64) {

	// Daily inflation
	inflationData := fetchDailyInflation(client)

	// Fetch new vesting claims
	vestingClaims := fetchVesting(client, currentBlock)

	// Compute data for UI
	computeInflation(inflationData, vestingClaims)
}

func fetchDailyInflation(client *ethclient.Client) interfaces.InflationData {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()

	inflationData := readInflation()

	// Check if we already got the inflation
	if len(inflationData.Inflation) > 0 && inflationData.Inflation[len(inflationData.Inflation)-1].Timestamp == startOfDay {
		return inflationData
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
	return inflationData
}

func getWeeklyFees() []interfaces.WeeklyFeesTable {
	return readWeeklyFees()
}

func fetchVesting(client *ethclient.Client, currentBlock uint64) []interfaces.VestingClaim {

	config := utils.ReadConfig(VESTING_CLAIMED_CONFIG_PATH)

	// veCRV
	from := config.LastBlock
	if from == 0 {
		from = 19809940 // Previous events was fetch from provider
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{
			utils.CRV_VESTING_CONTRACT,
		},
		Topics: [][]common.Hash{{common.HexToHash("0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	vestingClaims := readVestingClaims()

	for _, vLog := range logs {

		vestingContract, err := vesting.NewVesting(vLog.Address, client)
		if err != nil {
			panic(err)
		}

		event, err := vestingContract.ParseClaim(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(event.Raw.BlockNumber)))
		if err != nil {
			panic(err)
		}

		timestamp := block.Time()

		amountCRV := utils.Quo(event.Claimed, 18)
		tokenPrice := utils.GetHistoricalPriceTokenPrice(utils.CRV, "ethereum", timestamp)

		vestingClaims = append(vestingClaims, interfaces.VestingClaim{
			Timestamp: timestamp,
			AmountCRV: amountCRV,
			AmountUSD: amountCRV * tokenPrice,
		})

	}

	writeVestingClaims(vestingClaims)
	utils.WriteConfig(config, currentBlock, VESTING_CLAIMED_CONFIG_PATH)

	return vestingClaims
}

func computeInflation(inflationData interfaces.InflationData, vestingClaims []interfaces.VestingClaim) {
	weeklyFees := getWeeklyFees()

	// Group by month bounties + weekly fees vs inflation
	// Compute weekly fees
	inflationPerMonth := make([]interfaces.InflationPerMonth, 0)
	for _, weeklyFee := range weeklyFees {
		found := false
		ts := utils.GetStartOfMonth(int64(weeklyFee.Ts) / 1000)
		for i := range inflationPerMonth {
			if inflationPerMonth[i].Timestamp == ts {
				inflationPerMonth[i].WeeklyFeesIncomeAmountUSD += weeklyFee.Fees
				inflationPerMonth[i].TotalIncomeAmountUSD += weeklyFee.Fees
				found = true
				break
			}
		}

		if !found {
			inflationPerMonth = append(inflationPerMonth, interfaces.InflationPerMonth{
				Timestamp:                 ts,
				WeeklyFeesIncomeAmountUSD: weeklyFee.Fees,
				TotalIncomeAmountUSD:      weeklyFee.Fees,
			})
		}
	}

	// Bounties
	previousClaims := readDataPath(DATA_PATH)
	for _, claims := range previousClaims {
		found := false
		ts := utils.GetStartOfMonth(int64(claims.Timestamp))

		amountUSD := claims.AmountUSD
		for i := range inflationPerMonth {
			if inflationPerMonth[i].Timestamp == ts {
				inflationPerMonth[i].BountiesIncomeAmountUSD += amountUSD
				inflationPerMonth[i].TotalIncomeAmountUSD += amountUSD
				found = true
				break
			}
		}

		if !found {
			inflationPerMonth = append(inflationPerMonth, interfaces.InflationPerMonth{
				Timestamp:               ts,
				BountiesIncomeAmountUSD: amountUSD,
				TotalIncomeAmountUSD:    amountUSD,
			})
		}
	}

	// Inflation
	for _, inflation := range inflationData.Inflation {
		found := false
		ts := utils.GetStartOfMonth(int64(inflation.Timestamp))

		for i := range inflationPerMonth {
			if inflationPerMonth[i].Timestamp == ts {
				inflationPerMonth[i].InflationUSD += inflation.AmoutUSD
				inflationPerMonth[i].InflationCRV += inflation.AmountCRV
				found = true
				break
			}
		}

		if !found {
			inflationPerMonth = append(inflationPerMonth, interfaces.InflationPerMonth{
				Timestamp:    ts,
				InflationUSD: inflation.AmoutUSD,
				InflationCRV: inflation.AmountCRV,
			})
		}
	}

	// Vesting
	/*for _, vesting := range vestingClaims {
		found := false
		ts := utils.GetStartOfMonth(int64(vesting.Timestamp))

		for i := range inflationPerMonth {
			if inflationPerMonth[i].Timestamp == ts {
				inflationPerMonth[i].VestingUSD += vesting.AmountUSD
				inflationPerMonth[i].VestingCRV += vesting.AmountCRV
				inflationPerMonth[i].InflationUSD += vesting.AmountUSD
				inflationPerMonth[i].InflationCRV += vesting.AmountCRV
				found = true
				break
			}
		}

		if !found {
			inflationPerMonth = append(inflationPerMonth, interfaces.InflationPerMonth{
				Timestamp:    ts,
				VestingUSD:   vesting.AmountUSD,
				VestingCRV:   vesting.AmountCRV,
				InflationUSD: vesting.AmountUSD,
				InflationCRV: vesting.AmountCRV,
			})
		}
	}*/

	// Filter after 1st January 2023
	start := int64(1672527600) // 2023 January
	inflationPerMonthFiltered := make([]interfaces.InflationPerMonth, 0)
	for _, inflation := range inflationPerMonth {
		if inflation.Timestamp < start {
			continue
		}

		inflationPerMonthFiltered = append(inflationPerMonthFiltered, inflation)
	}

	sort.Slice(inflationPerMonthFiltered, func(i, j int) bool {
		return inflationPerMonthFiltered[i].Timestamp > inflationPerMonthFiltered[j].Timestamp
	})

	writeInflationPerMonth(inflationPerMonthFiltered)

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

func VestingHistorical(client *ethclient.Client, currentBlockf uint64, currentBlockTImestamp uint64) {
	if !utils.FileExists(VESTING_CLAIMED_PATH_TMP) {
		return
	}

	file, err := os.ReadFile(VESTING_CLAIMED_PATH_TMP)
	if err != nil {
		log.Fatal(err)
	}

	// Désérialiser le JSON dans une slice de VestingClaim.
	var claims [][]string
	err = json.Unmarshal(file, &claims)
	if err != nil {
		log.Fatal(err)
	}

	// Afficher les données pour vérification.
	vestingClaims := make([]interfaces.VestingClaim, 0)
	for _, claim := range claims {
		timestamp, err := strconv.Atoi(claim[0])
		if err != nil {
			log.Fatal(err)
		}
		amountBN, success := big.NewInt(0).SetString(claim[1], 10)
		if !success {
			log.Fatal("!bn")
		}

		amount := utils.Quo(amountBN, 18)

		tokenPrice := utils.GetHistoricalPriceTokenPrice(utils.CRV, "ethereum", uint64(timestamp))

		vestingClaims = append(vestingClaims, interfaces.VestingClaim{
			Timestamp: uint64(timestamp),
			AmountCRV: amount,
			AmountUSD: amount * tokenPrice,
		})
	}

	writeVestingClaims(vestingClaims)
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

func readVestingClaims() []interfaces.VestingClaim {

	if !utils.FileExists(VESTING_CLAIMED_PATH) {
		return make([]interfaces.VestingClaim, 0)
	}

	file, err := os.ReadFile(VESTING_CLAIMED_PATH)
	if err != nil {
		log.Fatal(err)
	}

	vestingClaims := make([]interfaces.VestingClaim, 0)
	if err := json.Unmarshal([]byte(file), &vestingClaims); err != nil {
		log.Fatal(err)
	}

	return vestingClaims
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

func writeInflationPerMonth(inflationPerMonthFiltered []interfaces.InflationPerMonth) {
	file, err := json.Marshal(inflationPerMonthFiltered)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(INFLATION_PER_MONTH_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func writeVestingClaims(vestingClaims []interfaces.VestingClaim) {
	file, err := json.Marshal(vestingClaims)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(VESTING_CLAIMED_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}
