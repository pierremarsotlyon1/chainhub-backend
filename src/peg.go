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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PEGS_PER_HOUR_PATH = "./data/pegs/pegs-per-hour.json"
)

var WRAPPERS = []interfaces.Wrapper{
	{
		Name:        "Stake DAO (sdCRV)",
		Address:     common.HexToAddress("0x52f541764E6e90eeBc5c21Ff570De0e2D63766B6"),
		LogoUrl:     "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fstakedao.42f95a34.png&w=48&q=75",
		PoolAddress: common.HexToAddress("0xCA0253A98D16e9C1e3614caFDA19318EE69772D0"),
	},
	{
		Name:        "Yearn (yCRV)",
		Address:     common.HexToAddress("0xF147b8125d2ef93FB6965Db97D6746952a133934"),
		LogoUrl:     "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fyearn.01d9002d.png&w=48&q=75",
		PoolAddress: common.HexToAddress("0x99f5acc8ec2da2bc0771c32814eff52b712de1e5"),
	},
	{
		Name:        "Convex (cvxCRV)",
		Address:     common.HexToAddress("0x989AEb4d175e16225E39E87d0D97A3360524AD80"),
		LogoUrl:     "https://www.defiwars.xyz/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fconvex.7542789f.jpeg&w=48&q=75",
		PoolAddress: common.HexToAddress("0x971add32ea87f10bd192671630be3be8a11b8623"),
	},
}

func Pegs(client *ethclient.Client) {

	pegs := readDailyPegs()

	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockTimestamp := head.Time
	startDay := uint64(utils.GetStartOfDay(blockTimestamp))

	for _, wrapper := range WRAPPERS {

		poolContract, err := curvePoolFactoryV2.NewCurvePoolFactoryV2(wrapper.PoolAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		opts := new(bind.CallOpts)
		opts.BlockNumber = head.Number

		pegBN, err := poolContract.GetDy(opts, big.NewInt(1), big.NewInt(0), utils.Mul(10000, 18))
		if err != nil {
			continue
		}

		peg := utils.Quo(pegBN, 18) / 100

		found := false
		for i := 0; i < len(pegs); i++ {
			if pegs[i].Timestamp == startDay && pegs[i].PoolAddress == wrapper.PoolAddress {
				pegs[i].Peg = peg
				found = true
				break
			}
		}

		if !found {
			pegs = append(pegs, interfaces.Peg{
				Timestamp:   startDay,
				PoolAddress: wrapper.PoolAddress,
				Peg:         peg,
			})
		}
	}

	writeDailyPegs(pegs)
}

func PegsHistorical(client *ethclient.Client) {
	startTimestamp := uint64(utils.GetStartOfDay(uint64(1686607200)))
	now := uint64(utils.GetStartOfDay(uint64(time.Now().Unix())))

	pegs := make([]interfaces.Peg, 0)
	for {

		blockNumber := utils.GetBlockNumberByTimestamp("ethereum", startTimestamp)

		for _, wrapper := range WRAPPERS {

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

	if !utils.FileExists(PEGS_PER_HOUR_PATH) {
		return make([]interfaces.Peg, 0)
	}

	file, err := os.ReadFile(PEGS_PER_HOUR_PATH)
	if err != nil {
		log.Fatal(err)
	}

	pegs := make([]interfaces.Peg, 0)
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
}
