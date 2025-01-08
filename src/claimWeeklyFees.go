package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/feeDistributor"
	"main/interfaces"
	"main/utils"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type feeDistributorConfig struct {
	Address        common.Address
	InceptionBlock uint64
	RewardSymbol   string
}

type claimWeeklyFee struct {
	TxHash         common.Hash `json:"txHash"`
	BlockTimestamp uint64      `json:"timestamp"`
	Amount         float64     `json:"amount"`
	RewardSymbol   string      `json:"symbol"`
}

const (
	claim_weekly_fees_config = "data/configs/claim-weekly-fees.json"
	bucket_all_claims_dir    = "data/weekly-fees-claims"
)

var feeDistributors = []feeDistributorConfig{
	feeDistributorConfig{
		Address:        utils.FEE_DISTRIBUTOR_THREE_CRV_MAINNET,
		InceptionBlock: 11278886,
		RewardSymbol:   "3CRV",
	},
	feeDistributorConfig{
		Address:        utils.FEE_DISTRIBUTOR_MAINNET,
		InceptionBlock: 20076216,
		RewardSymbol:   "crvUSD",
	},
}
var claim_topic = common.HexToHash("0x9cdcf2f7714cca3508c7f0110b04a90a80a3a8dd0e35de99689db74d28c5383e")

func ClaimWeeklyFees(client *ethclient.Client, currentBlock uint64) {
	config := utils.ReadConfig(claim_weekly_fees_config)

	for _, feeDistributor := range feeDistributors {
		fmt.Println("Fetching new events for distributor", feeDistributor.Address)
		newClaims := fetchClaimWeeklyFees(client, currentBlock, config, feeDistributor)
		fmt.Println("Found", len(newClaims), "claims")
		for user, claims := range newClaims {
			// Read old files
			userPath := bucket_all_claims_dir + "/" + strings.ToLower(user.Hex()) + ".json"
			fmt.Println("Writing user", userPath)
			oldUserClaims := make([]claimWeeklyFee, 0)
			fileData, err := utils.ReadBucketFile(userPath)
			if err == nil && len(fileData) > 0 {
				if err := json.Unmarshal(fileData, &oldUserClaims); err != nil {
					fmt.Println(err)
					continue
				}
			}

			// Append new events
			oldUserClaims = append(oldUserClaims, claims...)

			// Write the new file
			if err := utils.WriteBucketFile(userPath, oldUserClaims); err != nil {
				fmt.Println(err)
			}
		}
	}

	utils.WriteConfig(config, currentBlock, claim_weekly_fees_config)
}

func fetchClaimWeeklyFees(client *ethclient.Client, currentBlock uint64, config interfaces.Config, distributor feeDistributorConfig) map[common.Address][]claimWeeklyFee {
	from := config.LastBlock
	if from == 0 {
		from = distributor.InceptionBlock
	}

	run := true
	pas := int64(10000)
	claims := make(map[common.Address][]claimWeeklyFee)

	for run {
		to := int64(from) + pas
		if to > int64(currentBlock) {
			to = int64(currentBlock)
			run = false
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(from) + 1),
			ToBlock:   big.NewInt(to),
			Addresses: []common.Address{distributor.Address},
			Topics:    [][]common.Hash{{claim_topic}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			return claims
		}

		distributorContract, err := feeDistributor.NewFeeDistributor(distributor.Address, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Found", len(logs), "events from block", from, "to block", to)

		for _, vLog := range logs {

			block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
			if err != nil {
				fmt.Println(err)
				continue
			}

			event, err := distributorContract.ParseClaimed(vLog)
			if err != nil {
				fmt.Println(err)
				continue
			}

			_, exists := claims[event.Recipient]
			if !exists {
				claims[event.Recipient] = make([]claimWeeklyFee, 0)
			}

			claims[event.Recipient] = append(claims[event.Recipient], claimWeeklyFee{
				TxHash:         vLog.TxHash,
				BlockTimestamp: block.Time(),
				Amount:         utils.Quo(event.Amount, 18),
				RewardSymbol:   distributor.RewardSymbol,
			})
		}

		from += uint64(pas)
	}

	return claims
}
