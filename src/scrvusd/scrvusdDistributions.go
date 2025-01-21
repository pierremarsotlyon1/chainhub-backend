package scrvusd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/scrvUSD"
	"main/interfaces"
	"main/utils"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	scrvusd_distributions_config = "data/configs/scrvusd_distributions.json"
	bucket_scrvusd_distributions = "data/scrvusd/distributions.json"
)

func ScrvusdDistributions(client *ethclient.Client, currentBlock uint64) {

	// Read the config
	config := utils.ReadConfig(scrvusd_distributions_config)

	// Read last token exchanges
	distributions := readScrvUSDDistributions()

	// Fetch new ones
	distributions = append(distributions, fetchDistributions(client, config, currentBlock)...)

	// Sort by DESC
	sort.Slice(distributions, func(i, j int) bool { return distributions[i].BlockTimestamp > distributions[j].BlockTimestamp })

	// Write files
	utils.WriteConfig(config, currentBlock, scrvusd_distributions_config)
	writeScrvUSDDistributions(distributions)
}

func fetchDistributions(client *ethclient.Client, config interfaces.Config, currentBlock uint64) []interfaces.ScrvusdDistribution {
	from := config.LastBlock
	if from == 0 {
		from = 21087889
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from) + 1),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{utils.SCRV_USD},
		Topics:    [][]common.Hash{{common.HexToHash("0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	distributions := make([]interfaces.ScrvusdDistribution, 0)
	if err != nil {
		return distributions
	}

	scrvUSDContract, err := scrvUSD.NewScrvUSD(utils.SCRV_USD, client)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		event, err := scrvUSDContract.ParseStrategyReported(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		gain := utils.Quo(event.Gain, 18)

		crvUSDPrice := utils.GetHistoricalPriceTokenPrice(utils.CRVUSD_ADDRESS, "ethereum", block.Time())
		gainUSD := gain * crvUSDPrice

		distributions = append(distributions, interfaces.ScrvusdDistribution{
			TxHash:         vLog.TxHash,
			BlockTimestamp: block.Time(),
			Gain:           gain,
			GainUSD:        gainUSD,
		})

	}

	return distributions
}

func readScrvUSDDistributions() []interfaces.ScrvusdDistribution {
	distributions := make([]interfaces.ScrvusdDistribution, 0)
	b, err := utils.ReadBucketFile(bucket_scrvusd_distributions)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &distributions); err != nil {
			log.Fatal(err)
		}
		return distributions
	}
	return distributions
}

func writeScrvUSDDistributions(distributions []interfaces.ScrvusdDistribution) {
	if err := utils.WriteBucketFile(bucket_scrvusd_distributions, distributions); err != nil {
		fmt.Println(err)
	}
}
