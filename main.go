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

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("No arg")
	}

	functionToExecute := argsWithoutProg[0]

	ALCHEMY_RPC_URL = utils.GetPublicRpcUrl("mainnet")

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
		clientDrpc, err := ethclient.Dial(utils.GetPublicRpcUrl("mainnet"))
		if err != nil {
			panic(err)
		}
		src.PendingPoolFees(clientDrpc, lastestBlock.NumberU64(), lastestBlock.Time())
	case "inflation":
		src.Inflation(client, lastestBlock.NumberU64(), lastestBlock.Time())
	case "llamalend":
		src.LlamalendHistorical()
		src.LlamalendFixSort()
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
	case "checkpoint":
		src.CheckpointGauges()
	}
}
