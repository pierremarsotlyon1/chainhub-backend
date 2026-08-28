package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/curvePoolFactoryV2"
	"main/contracts/multicall"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PEGS_PER_HOUR_PATH = "./data/pegs/pegs-per-hour.json"

	BUCKET_PEGS_DIR           = "data/pegs"
	BUCKET_PEGS_PER_HOUR_FILE = BUCKET_PEGS_DIR + "/pegs-per-hour.json"

	PEG_FETCH_WORKERS = 8
)

func Pegs(client *ethclient.Client) {

	pegs := readDailyPegs()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockTimestamp := block.Time()
	currentDay := uint64(utils.GetStartOfDay(blockTimestamp))

	// Backfill missing days since the last timestamp in the file
	lastDay := currentDay
	if len(pegs) > 0 {
		lastDay = pegs[0].Timestamp
		for _, peg := range pegs {
			if peg.Timestamp > lastDay {
				lastDay = peg.Timestamp
			}
		}
	}

	days := make([]uint64, 0)
	for day := lastDay + utils.DAY_TO_SEC; day < currentDay; day += utils.DAY_TO_SEC {
		days = append(days, day)
	}

	log.Println("Pegs: backfilling", len(days), "missing days")
	pegs = upsertPegs(pegs, fetchDaysPegs(client, days))

	log.Println("Pegs: fetching current day at block", block.Number())
	pegs = upsertPegs(pegs, fetchPegsAtBlock(client, currentDay, block.Number()))

	writeDailyPegs(pegs)
	log.Println("Pegs: done,", len(pegs), "entries written")
}

func PegsHistorical(client *ethclient.Client) {
	startTimestamp := uint64(utils.GetStartOfDay(uint64(1686607200)))
	now := uint64(utils.GetStartOfDay(uint64(time.Now().Unix())))

	days := make([]uint64, 0)
	for day := startTimestamp; day <= now; day += utils.DAY_TO_SEC {
		days = append(days, day)
	}

	log.Println("PegsHistorical: fetching", len(days), "days")
	pegs := fetchDaysPegs(client, days)

	writeDailyPegs(pegs)
	log.Println("PegsHistorical: done,", len(pegs), "entries written")
}

// fetchDaysPegs fetches the pegs of all wrappers for each day concurrently
func fetchDaysPegs(client *ethclient.Client, days []uint64) []interfaces.Peg {
	results := make([][]interfaces.Peg, len(days))

	var wg sync.WaitGroup
	var fetched atomic.Uint64
	sem := make(chan struct{}, PEG_FETCH_WORKERS)

	for i, day := range days {
		wg.Add(1)
		go func(i int, day uint64) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			blockNumber := utils.GetBlockNumberByTimestamp("ethereum", day)
			if blockNumber == 0 {
				fmt.Println("No block number found for timestamp", day)
				return
			}

			results[i] = fetchPegsAtBlock(client, day, big.NewInt(int64(blockNumber)))
			log.Println("Pegs:", fetched.Add(1), "/", len(days), "days fetched (day", day, ", block", blockNumber, ")")
		}(i, day)
	}

	wg.Wait()

	pegs := make([]interfaces.Peg, 0, len(days)*len(utils.WRAPPERS))
	for _, dayPegs := range results {
		pegs = append(pegs, dayPegs...)
	}

	return pegs
}

// fetchPegsAtBlock fetches the pegs of all wrappers in a single multicall
func fetchPegsAtBlock(client *ethclient.Client, day uint64, blockNumber *big.Int) []interfaces.Peg {
	poolABI, err := curvePoolFactoryV2.CurvePoolFactoryV2MetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	callData, err := poolABI.Pack("get_dy", big.NewInt(1), big.NewInt(0), utils.Mul(10000, 18))
	if err != nil {
		log.Fatal(err)
	}

	calls := make([]multicall.Multicall3Call, 0, len(utils.WRAPPERS))
	for _, wrapper := range utils.WRAPPERS {
		calls = append(calls, multicall.Multicall3Call{
			Target:   wrapper.PoolAddress,
			CallData: callData,
		})
	}

	var responses []utils.MulticallResponse
	for attempt := 1; attempt <= 3; attempt++ {
		responses = utils.Multicall(client, calls, utils.MULTICALL_MAINNET, blockNumber)
		if len(responses) == len(calls) {
			break
		}
		time.Sleep(time.Duration(attempt) * time.Second)
	}

	if len(responses) != len(calls) {
		fmt.Println("Multicall failed for block", blockNumber)
		return nil
	}

	pegs := make([]interfaces.Peg, 0, len(utils.WRAPPERS))
	for i, response := range responses {
		if !response.Success || len(response.ReturnData) == 0 {
			continue
		}

		values, err := poolABI.Unpack("get_dy", response.ReturnData)
		if err != nil || len(values) == 0 {
			fmt.Println(err, utils.WRAPPERS[i].PoolAddress, blockNumber)
			continue
		}

		pegBN, ok := values[0].(*big.Int)
		if !ok {
			continue
		}

		pegs = append(pegs, interfaces.Peg{
			Timestamp:   day,
			PoolAddress: utils.WRAPPERS[i].PoolAddress,
			Peg:         utils.Quo(pegBN, 18) / 100,
		})
	}

	return pegs
}

func upsertPegs(pegs []interfaces.Peg, newPegs []interfaces.Peg) []interfaces.Peg {
	for _, newPeg := range newPegs {
		found := false
		for i := 0; i < len(pegs); i++ {
			if pegs[i].Timestamp == newPeg.Timestamp && pegs[i].PoolAddress == newPeg.PoolAddress {
				pegs[i].Peg = newPeg.Peg
				found = true
				break
			}
		}

		if !found {
			pegs = append(pegs, newPeg)
		}
	}

	return pegs
}

func readDailyPegs() []interfaces.Peg {

	pegs := make([]interfaces.Peg, 0)
	b, err := utils.ReadBucketFile(BUCKET_PEGS_PER_HOUR_FILE)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &pegs); err != nil {
			log.Fatal(err)
		}
		return pegs
	}

	if !utils.FileExists(PEGS_PER_HOUR_PATH) {
		return make([]interfaces.Peg, 0)
	}

	file, err := os.ReadFile(PEGS_PER_HOUR_PATH)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &pegs); err != nil {
		log.Fatal(err)
	}

	return pegs
}

func writeDailyPegs(pegs []interfaces.Peg) {
	file, err := json.Marshal(pegs)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PEGS_PER_HOUR_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}

	if err := utils.WriteBucketFile(BUCKET_PEGS_PER_HOUR_FILE, pegs); err != nil {
		fmt.Println(err)
	}
}
