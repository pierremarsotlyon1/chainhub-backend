package main

import (
	"context"
	"fmt"
	"log"
	"main/src"
	"main/src/scrvusd"
	"main/utils"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ALCHEMY_RPC_URL = ""
var RPC_LOCAL_NODE = "/root/geth/data/geth/geth.ipc"

// First arg : script to execut
// Second arg : alchemy api key, if not set, get default one
func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("No arg")
	}

	functionToExecute := argsWithoutProg[0]

	alchemyApiKey := utils.GoDotEnvVariable("DRPC_API_KEY")

	if len(argsWithoutProg) > 1 {
		alchemyApiKey = utils.GoDotEnvVariable(argsWithoutProg[1])
	}

	if len(alchemyApiKey) == 0 {
		panic("ALCHEMY_APIKEY not set")
	}

	ALCHEMY_RPC_URL = "https://lb.drpc.org/ethereum/" + alchemyApiKey

	client, err := ethclient.Dial(RPC_LOCAL_NODE)
	if err != nil {
		client, err = ethclient.Dial(ALCHEMY_RPC_URL)
		if err != nil {
			panic(err)
		}
	}

	lastestBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Arg : ", functionToExecute)

	switch functionToExecute {
	case "bounties":
		src.FetchBounties(client, lastestBlock.NumberU64(), ALCHEMY_RPC_URL, utils.GoDotEnvVariable("DRPC_API_KEY"))
	case "locks":
		src.FetchLocks(client, lastestBlock.NumberU64())
	case "votes":
		src.FetchVotes(client, lastestBlock.NumberU64())
	case "gauges":
		src.FetchGaugeWeights(ALCHEMY_RPC_URL)
	case "crvusd":
		src.CrvUSD(client, lastestBlock.NumberU64())
	case "pegkeepers":
		src.FetchPegKeepers(client, lastestBlock.NumberU64(), lastestBlock.Time())
	case "pendingPoolFees":
		clientDrpc, err := ethclient.Dial(utils.GoDotEnvVariable("DRPC_MAINNET"))
		if err != nil {
			panic(err)
		}
		src.PendingPoolFees(clientDrpc, lastestBlock.NumberU64(), lastestBlock.Time())
	case "inflation":
		src.Inflation(client, lastestBlock.NumberU64(), lastestBlock.Time())
	case "llamalend":
		src.Llamalend()
	case "peg":
		src.Pegs(client)
	case "weeklyFees":
		src.WeeklyFees(client)
	case "wrappersSwaps":
		src.WrappersSwap(client, lastestBlock.NumberU64())
	case "claimWeeklyFees":
		src.ClaimWeeklyFees(client, lastestBlock.NumberU64())
	case "scrvusdDistributions":
		scrvusd.ScrvusdDistributions(client, lastestBlock.NumberU64())
	case "scrvusdApys":
		scrvusd.ScrvusdApys(client, lastestBlock.NumberU64())
	case "curveApiCache":
		src.CurveApiCache()
	}
}
