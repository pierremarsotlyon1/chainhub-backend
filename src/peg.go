package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/curvePoolFactoryV2"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PEGS_PER_HOUR_PATH = "./data/pegs/pegs-per-hour.json"

	BUCKET_PEGS_DIR           = "data/pegs"
	BUCKET_PEGS_PER_HOUR_FILE = BUCKET_PEGS_DIR + "/pegs-per-hour.json"
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

	for day := lastDay + utils.DAY_TO_SEC; day < currentDay; day += utils.DAY_TO_SEC {
		blockNumber := utils.GetBlockNumberByTimestamp("ethereum", day)
		pegs = fetchPegsAtBlock(client, pegs, day, big.NewInt(int64(blockNumber)))
	}

	pegs = fetchPegsAtBlock(client, pegs, currentDay, block.Number())

	writeDailyPegs(pegs)
}

func fetchPegsAtBlock(client *ethclient.Client, pegs []interfaces.Peg, day uint64, blockNumber *big.Int) []interfaces.Peg {

	for _, wrapper := range utils.WRAPPERS {

		poolContract, err := curvePoolFactoryV2.NewCurvePoolFactoryV2(wrapper.PoolAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		opts := new(bind.CallOpts)
		opts.BlockNumber = blockNumber

		pegBN, err := poolContract.GetDy(opts, big.NewInt(1), big.NewInt(0), utils.Mul(10000, 18))
		if err != nil {
			continue
		}

		peg := utils.Quo(pegBN, 18) / 100

		found := false
		for i := 0; i < len(pegs); i++ {
			if pegs[i].Timestamp == day && pegs[i].PoolAddress == wrapper.PoolAddress {
				pegs[i].Peg = peg
				found = true
				break
			}
		}

		if !found {
			pegs = append(pegs, interfaces.Peg{
				Timestamp:   day,
				PoolAddress: wrapper.PoolAddress,
				Peg:         peg,
			})
		}
	}

	return pegs
}

func PegsHistorical(client *ethclient.Client) {
	startTimestamp := uint64(utils.GetStartOfDay(uint64(1686607200)))
	now := uint64(utils.GetStartOfDay(uint64(time.Now().Unix())))

	pegs := make([]interfaces.Peg, 0)
	for {

		blockNumber := utils.GetBlockNumberByTimestamp("ethereum", startTimestamp)

		for _, wrapper := range utils.WRAPPERS {

			poolContract, err := curvePoolFactoryV2.NewCurvePoolFactoryV2(wrapper.PoolAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			opts := new(bind.CallOpts)
			opts.BlockNumber = big.NewInt(int64(blockNumber))

			pegBN, err := poolContract.GetDy(opts, big.NewInt(1), big.NewInt(0), utils.Mul(10000, 18))
			if err != nil {
				fmt.Println(err, wrapper.PoolAddress, blockNumber)
				continue
			}

			pegs = append(pegs, interfaces.Peg{
				Timestamp:   startTimestamp,
				PoolAddress: wrapper.PoolAddress,
				Peg:         utils.Quo(pegBN, 18) / 100,
			})
		}

		startTimestamp += utils.DAY_TO_SEC

		if startTimestamp > now {
			break
		}
	}

	writeDailyPegs(pegs)
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
