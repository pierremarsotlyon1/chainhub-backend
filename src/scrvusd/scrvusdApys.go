package scrvusd

import (
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/scrvUSD"
	"main/interfaces"
	"main/utils"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	bucket_scrvusd_apys = "data/scrvusd/apys.json"
	default_timestamp   = 1730418191
)

func ScrvusdApys(client *ethclient.Client, currentBlock uint64) {

	// Read last apys
	apys := readScrvUSDApys()

	// Fetch new ones
	apys = fetchApys(client, apys)

	// Sort by DESC
	sort.Slice(apys, func(i, j int) bool { return apys[i].Ts > apys[j].Ts })

	writeScrvUSDApys(apys)
}

func fetchApys(client *ethclient.Client, apys []interfaces.ScrvusdApy) []interfaces.ScrvusdApy {

	// Get current start of day
	now := time.Now().Unix()
	startDay := utils.GetStartOfDay(uint64(now))

	// Fetch historical apys
	// Get last timestamp fetched
	lastTimestamp := default_timestamp

	if len(apys) > 0 {
		lastTimestamp = int(apys[0].Ts)
	}

	lastTimestampStartDay := utils.GetStartOfDay(uint64(lastTimestamp))

	for lastTimestampStartDay < startDay {
		apy := getApy(client, lastTimestampStartDay)

		apys = append(apys, interfaces.ScrvusdApy{
			Apy: apy,
			Ts:  uint64(lastTimestampStartDay),
		})

		lastTimestampStartDay += int64(utils.DAY_TO_SEC)
	}

	// Sort by timestamp DESC
	sort.Slice(apys, func(i, j int) bool { return apys[i].Ts > apys[j].Ts })

	// Get current apy
	apy := getApy(client, startDay)

	if len(apys) > 0 && apys[0].Ts == uint64(startDay) {
		apys[0].Apy = apy
	} else {
		apys = append(apys, interfaces.ScrvusdApy{
			Apy: apy,
			Ts:  uint64(lastTimestampStartDay),
		})
	}

	return apys
}

func getApy(client *ethclient.Client, timestamp int64) float64 {
	block := utils.GetBlockNumberByTimestamp("ethereum", uint64(timestamp))

	scrvContract, err := scrvUSD.NewScrvUSD(utils.SCRV_USD, client)
	if err != nil {
		fmt.Println("toto", err)
		return 0
	}

	opts := bind.CallOpts{}
	opts.BlockNumber = big.NewInt(int64(block))

	rateBN, err := scrvContract.ProfitUnlockingRate(&opts)
	if err != nil {
		fmt.Println("ProfitUnlockingRate", block, err)
		return 0
	}

	totalSupplyBN, err := scrvContract.TotalSupply(&opts)
	if err != nil {
		fmt.Println("TotalSupply", err)
		return 0
	}

	rate := utils.Quo(rateBN, 18)
	totalSupply := utils.Quo(totalSupplyBN, 18)

	if totalSupply < 1 {
		totalSupply = 1
	}

	return ((rate / 1e12) * 31536000 * 100) / totalSupply
}

func readScrvUSDApys() []interfaces.ScrvusdApy {
	apys := make([]interfaces.ScrvusdApy, 0)
	b, err := utils.ReadBucketFile(bucket_scrvusd_apys)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &apys); err != nil {
			log.Fatal(err)
		}
		return apys
	}
	return apys
}

func writeScrvUSDApys(apys []interfaces.ScrvusdApy) {
	if err := utils.WriteBucketFile(bucket_scrvusd_apys, apys); err != nil {
		fmt.Println(err)
	}
}
