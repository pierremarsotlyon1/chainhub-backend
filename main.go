package main

import (
	"context"
	"fmt"
	"log"
	"main/src"
	"main/utils"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var ALCHEMY_RPC_URL = ""
var RPC_LOCAL_NODE = "/datastore/.ethereum/geth.ipc"

// First arg : script to execut
// Second arg : alchemy api key, if not set, get default one
func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("No arg")
	}

	functionToExecute := argsWithoutProg[0]

	alchemyApiKey := utils.GoDotEnvVariable("ALCHEMY_APIKEY")

	if len(argsWithoutProg) > 1 {
		alchemyApiKey = utils.GoDotEnvVariable(argsWithoutProg[1])
	}

	if len(alchemyApiKey) == 0 {
		panic("ALCHEMY_APIKEY not set")
	}

	ALCHEMY_RPC_URL = "https://eth-mainnet.g.alchemy.com/v2/" + alchemyApiKey

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
		src.FetchBounties(client, lastestBlock.NumberU64(), ALCHEMY_RPC_URL)
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
		src.PendingPoolFees(client, lastestBlock.NumberU64(), lastestBlock.Time())
	case "inflation":
		src.Inflation(client, lastestBlock.NumberU64(), lastestBlock.Time())
	case "llamalend":
		src.Llamalend()
	case "peg":
		src.Pegs(client)
	}
}
