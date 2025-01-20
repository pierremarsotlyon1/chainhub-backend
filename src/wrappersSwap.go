package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/curvePoolWrappers"
	"main/contracts/erc20"
	"main/interfaces"
	"main/utils"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	wrappers_config       = "data/configs/wrappers_swap.json"
	bucket_wrappers_swaps = "data/wrappers/swaps.json"
	max_swaps             = 100
)

func WrappersSwap(client *ethclient.Client, currentBlock uint64) {

	// Read the config
	config := utils.ReadConfig(wrappers_config)

	// Read last token exchanges
	wrappersSwaps := readWrappersSwaps()

	// Fetch new events
	for _, wrapper := range utils.WRAPPERS {
		wrappersSwaps = append(wrappersSwaps, fetchTokenExchanges(client, config, currentBlock, wrapper)...)
	}

	// Sort swaps DESC
	sort.Slice(wrappersSwaps, func(i, j int) bool { return wrappersSwaps[i].BlockTimestamp > wrappersSwaps[j].BlockTimestamp })

	// Keep max 100 swaps
	if len(wrappersSwaps) > max_swaps {
		wrappersSwaps = wrappersSwaps[0:max_swaps]
	}
	// Write files
	utils.WriteConfig(config, currentBlock, wrappers_config)
	writeWrappersSwaps(wrappersSwaps)
}

func fetchTokenExchanges(client *ethclient.Client, config interfaces.Config, currentBlock uint64, wrapper interfaces.Wrapper) []interfaces.PoolTokenExchange {
	from := config.LastBlock
	if from == 0 {
		from = 21528043
	}

	pools := []common.Address{wrapper.PoolAddress}
	pools = append(pools, wrapper.OldPoolAddresses...)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from) + 1),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: pools,
		Topics:    [][]common.Hash{{common.HexToHash("0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	swaps := make([]interfaces.PoolTokenExchange, 0)
	if err != nil {
		return swaps
	}

	poolContract, err := curvePoolWrappers.NewCurvePoolWrappers(wrapper.PoolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		event, err := poolContract.ParseTokenExchange(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Token sold
		tokenSoldAddress, err := poolContract.Coins(nil, event.SoldId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tokenSoldAmount := utils.Quo(event.TokensSold, 18)

		tokenSoldContract, err := erc20.NewErc20(tokenSoldAddress, client)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tokenSoldSymbol, err := tokenSoldContract.Symbol(nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Token bought
		tokenBoughtAddress, err := poolContract.Coins(nil, event.BoughtId)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tokenBoughtAmount := utils.Quo(event.TokensBought, 18)

		tokenBoughtContract, err := erc20.NewErc20(tokenBoughtAddress, client)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tokenBoughtSymbol, err := tokenBoughtContract.Symbol(nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		swaps = append(swaps, interfaces.PoolTokenExchange{
			User:               common.HexToAddress("0x" + vLog.Topics[1].Hex()[26:]),
			TokenSoldAddress:   tokenSoldAddress,
			TokenSoldAmount:    tokenSoldAmount,
			TokenBoughtAddress: tokenBoughtAddress,
			TokenBoughtAmount:  tokenBoughtAmount,
			SymbolSold:         tokenSoldSymbol,
			SymbolBought:       tokenBoughtSymbol,
			TxHash:             vLog.TxHash,
			BlockTimestamp:     block.Time(),
			PoolAddress:        vLog.Address,
		})

	}

	return swaps
}

func readWrappersSwaps() []interfaces.PoolTokenExchange {
	wrappersSwaps := make([]interfaces.PoolTokenExchange, 0)
	b, err := utils.ReadBucketFile(bucket_wrappers_swaps)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &wrappersSwaps); err != nil {
			log.Fatal(err)
		}
		return wrappersSwaps
	}
	return wrappersSwaps
}

func writeWrappersSwaps(swaps []interfaces.PoolTokenExchange) {
	// Remove duplicate
	newSwaps := make([]interfaces.PoolTokenExchange, 0)
	for _, swap := range swaps {
		found := false
		for _, newSwap := range newSwaps {
			if swap.TxHash.Hex() == newSwap.TxHash.Hex() {
				found = true
				break
			}
		}

		if !found {
			newSwaps = append(newSwaps, swap)
		}
	}

	if err := utils.WriteBucketFile(bucket_wrappers_swaps, newSwaps); err != nil {
		fmt.Println(err)
	}
}
