package src

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/erc20"
	"main/contracts/questDistributor"
	"main/contracts/vmRemoteManager"
	"main/contracts/vmTokenFactory"
	"main/contracts/vmV2"
	"main/contracts/votemarketV1"
	"main/contracts/votemarketV2"
	"main/contracts/votiumMerkle"
	"main/contracts/votiumVECrv"
	"main/contracts/yBribeV3"
	"main/interfaces"
	"main/utils"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var VOTIUM_VE_CRV_ADDRESSES = []common.Address{
	common.HexToAddress("0xB4Fb1FD4AEC780BC255bF231189E9A244475d260"),
}

var VOTIUM_MERKLE_V2 = []common.Address{
	common.HexToAddress("0x378Ba9B73309bE80BF4C2c027aAD799766a7ED5A"),
}

var VOTEMARKET_ADDRESSES_V1 = []common.Address{
	common.HexToAddress("0x7D0F747eb583D43D41897994c983F13eF7459e1f"),
}

var VOTEMARKET_ADDRESSES_V2 = []common.Address{
	common.HexToAddress("0x0000000BE1d98523B5469AfF51A1e7b4891c6225"),
	common.HexToAddress("0x0000000895cB182E6f983eb4D8b4E0Aa0B31Ae4c"),
	common.HexToAddress("0x000000073D065Fc33a3050C2d0E19C393a5699ba"),
}

var VOTEMARKET_VMV2_ADDRESSES_V1 = []interfaces.VmV2Config{
	{
		ChainId:            42161,
		DefilammaChainName: "arbitrum",
		DrpcChainName:      "arbitrum",
		Markets: []interfaces.VmMarket{
			{
				Addresses: []common.Address{
					common.HexToAddress("0x5e5C922a5Eeab508486eB906ebE7bDFFB05D81e5"),
				},
				BlockNumberDeployed: 269278364,
			},
			{
				Addresses: []common.Address{
					common.HexToAddress("0x8c2c5A295450DDFf4CB360cA73FCCC12243D14D9"),
				},
				BlockNumberDeployed: 306978544,
			},
		},
	},
	{
		ChainId:            10,
		DefilammaChainName: "optimism",
		DrpcChainName:      "optimism",
		Markets: []interfaces.VmMarket{
			{
				Addresses: []common.Address{
					common.HexToAddress("0x5e5C922a5Eeab508486eB906ebE7bDFFB05D81e5"),
				},
				BlockNumberDeployed: 127353052,
			},
			{
				Addresses: []common.Address{
					common.HexToAddress("0x8c2c5A295450DDFf4CB360cA73FCCC12243D14D9"),
				},
				BlockNumberDeployed: 132095453,
			},
		},
	},
	{
		ChainId:            137,
		DefilammaChainName: "polygon",
		DrpcChainName:      "polygon",
		Markets: []interfaces.VmMarket{
			{
				Addresses: []common.Address{
					common.HexToAddress("0x5e5C922a5Eeab508486eB906ebE7bDFFB05D81e5"),
				},
				BlockNumberDeployed: 63679441,
			},
			{
				Addresses: []common.Address{
					common.HexToAddress("0x8c2c5A295450DDFf4CB360cA73FCCC12243D14D9"),
				},
				BlockNumberDeployed: 68030935,
			},
		},
	},
	{
		ChainId:            8453,
		DefilammaChainName: "base",
		DrpcChainName:      "base",
		Markets: []interfaces.VmMarket{
			{
				Addresses: []common.Address{
					common.HexToAddress("0x5e5C922a5Eeab508486eB906ebE7bDFFB05D81e5"),
				},
				BlockNumberDeployed: 21757790,
			},
			{
				Addresses: []common.Address{
					common.HexToAddress("0x8c2c5A295450DDFf4CB360cA73FCCC12243D14D9"),
				},
				BlockNumberDeployed: 26500169,
			},
		},
	},
}

var QUEST_ADDRESSES = []common.Address{
	common.HexToAddress("0x3682518b529e4404fb05250F9ad590C3218E5F9f"),
	common.HexToAddress("0xce6dc32252d85e2e955Bfd3b85660917F040a933"),
	common.HexToAddress("0x8EdcFE9Bc7d2a735117B94C16456D8303777abbb"),
	common.HexToAddress("0x358549D4Cb7f97f389812B86673a6cf8c1FF59D2"),
	common.HexToAddress("0x999881aA210B637ffF7d22c8566319444B38695B"),
	common.HexToAddress("0x1327C85CE6F3C83faBC4f5C294F57ac05bCb51eB"),
}

var YBRIBE_ADDRESSES = []common.Address{
	common.HexToAddress("0x03dFdBcD4056E2F92251c7B07423E1a33a7D3F6d"),
}

var BRIBE_CRV_FINANCE_ADDRESSES = []common.Address{
	common.HexToAddress("0x7893bbb46613d7a4fbcc31dab4c9b823ffee1026"),
}

const (
	data_path       = "./data/data.json"
	stats_path      = "./data/stats.json"
	bounties_config = "./data/configs/bounties-config.json"

	BUCKET_BOUNTIES_DIR        = "data/bounties"
	BUCKET_BOUNTIES_DATA_FILE  = BUCKET_BOUNTIES_DIR + "/data.json"
	BUCKET_BOUNTIES_STATS_FILE = BUCKET_BOUNTIES_DIR + "/stats.json"
)

func FetchBounties(client *ethclient.Client, currentBlock uint64, alchemyRpcUrl string, drpcKey string) {

	// Fetch curve pools
	curvePools, _ := utils.GetAllCurvePools()

	config := utils.ReadConfig(bounties_config)

	// Bounties
	allClaimed := make([]interfaces.BountyClaimed, 0)

	previousClaims := readDataPath()

	allClaimed = append(allClaimed, previousClaims...)

	fmt.Println("Fetching votium")
	allClaimed = append(allClaimed, fetchVotium(client, curvePools, currentBlock, config)...)

	fmt.Println("Fetching votemarket v1")
	allClaimed = append(allClaimed, fetchVotemarketV1(client, curvePools, currentBlock, config)...)

	fmt.Println("Fetching votemarket v2")
	allClaimed = append(allClaimed, fetchVotemarketV2(client, curvePools, currentBlock, config)...)

	fmt.Println("Fetching votemarket vm v2")
	allClaimed = append(allClaimed, fetchVotemarketVmV2V1(curvePools, drpcKey)...)

	fmt.Println("Fetching quest")
	allClaimed = append(allClaimed, fetchQuest(client, curvePools, currentBlock, config)...)

	fmt.Println("Fetching yBribe")
	allClaimed = append(allClaimed, fetchYBribe(client, curvePools, currentBlock, config, alchemyRpcUrl)...)

	fmt.Println("Fetching bribe crv finance")
	allClaimed = append(allClaimed, fetchBribeCrvFinance(client, curvePools, currentBlock, config, alchemyRpcUrl)...)

	fmt.Println(len(allClaimed), " claims found")

	// Compute stats files
	computeBountiesStats(allClaimed)

	// Write config
	utils.WriteConfig(config, currentBlock, bounties_config)
}

func computeBountiesStats(allClaimed []interfaces.BountyClaimed) {
	totalBountiesUSD := 0.0
	totalVeCRVBountiesUSD := 0.0
	totalVlCVXBountiesUSD := 0.0

	allClaimedWithPrice := make([]interfaces.BountyClaimed, 0)
	now := uint64(time.Now().Unix())
	for _, bounty := range allClaimed {

		totalBountiesUSD += bounty.AmountUSD
		if isVlCVX(bounty) {
			totalVlCVXBountiesUSD += bounty.AmountUSD
		} else {
			totalVeCRVBountiesUSD += bounty.AmountUSD
		}

		allClaimedWithPrice = append(allClaimedWithPrice, bounty)
	}

	writeDataPath(allClaimedWithPrice)

	fmt.Println("Total bounties USD (veCRV + vlCVX): ", fmt.Sprintf("%f", totalBountiesUSD))
	fmt.Println("Total veCRV bounties USD : ", fmt.Sprintf("%f", totalVeCRVBountiesUSD))
	fmt.Println("Total vlCVX bounties USD : ", fmt.Sprintf("%f", totalVlCVXBountiesUSD))

	// Sort claims by timestamp ASC
	sort.Slice(allClaimedWithPrice, func(i, j int) bool { return allClaimedWithPrice[i].Timestamp < allClaimedWithPrice[j].Timestamp })

	// Remove future bounties
	var allClaimedFiltered []interfaces.BountyClaimed
	for _, b := range allClaimedWithPrice {
		if b.Timestamp < now {
			allClaimedFiltered = append(allClaimedFiltered, b)
		}
	}

	timestampEnd := allClaimedFiltered[len(allClaimedFiltered)-1].Timestamp
	timestampStart := allClaimedFiltered[0].Timestamp

	var stats interfaces.Stats

	stats.TotalClaimed = totalBountiesUSD
	stats.VeCRVTotalClaimed = totalVeCRVBountiesUSD
	stats.VlCVXTotalClaimed = totalVlCVXBountiesUSD

	stats.ClaimsLast7Days = generateDaysData(allClaimedFiltered, timestampEnd-7*utils.DAY_TO_SEC, 25*utils.MIN_TO_SEC)
	stats.ClaimsLast30Days = generateDaysData(allClaimedFiltered, timestampEnd-30*utils.DAY_TO_SEC, 2*utils.HOUR_TO_SEC)
	stats.ClaimsLast180Days = generateDaysData(allClaimedFiltered, timestampEnd-180*utils.DAY_TO_SEC, 12*utils.HOUR_TO_SEC)
	stats.ClaimsLast365Days = generateDaysData(allClaimedFiltered, timestampEnd-365*utils.DAY_TO_SEC, utils.DAY_TO_SEC)
	stats.ClaimsSinceInception = generateDaysData(allClaimedFiltered, timestampStart, 15*utils.DAY_TO_SEC)

	writeStats(stats)
}
func fetchVotium(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config) []interfaces.BountyClaimed {
	fromVeCRV := config.LastBlock
	if fromVeCRV == 0 {
		fromVeCRV = 14730003
	}

	const maxBlockRange = 499
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)

	// veCRV
	for start := fromVeCRV + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := min(start+maxBlockRange, currentBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: VOTIUM_VE_CRV_ADDRESSES,
			Topics:    [][]common.Hash{{common.HexToHash("0x51c8cd367a987b8c2f652c101ea7076ec8e4dfd33c4c77bb80e018e7143b6512")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, vLog := range logs {
			votiumContract, err := votiumVECrv.NewVotiumVECrv(vLog.Address, client)
			if err != nil {
				fmt.Println(err)
				continue
			}

			event, err := votiumContract.ParseNewReward(vLog)
			if err != nil {
				fmt.Println(err)
				continue
			}

			bountiesClaimed = addVotiumClaim(client, bountiesClaimed, vLog.Address, event.Token, vLog.BlockNumber, event.Amount, vLog.TxHash, "")
		}
	}

	// vlCVX V2
	fromVLCVX := config.LastBlock
	if fromVLCVX == 0 {
		fromVLCVX = 13320168
	}

	for start := fromVLCVX + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := start + maxBlockRange
		if end > currentBlock {
			end = currentBlock
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: VOTIUM_MERKLE_V2,
			Topics:    [][]common.Hash{{common.HexToHash("0x4766921f5c59646d22d7d266a29164c8e9623684d8dfdbd931731dfdca025238")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, vLog := range logs {
			votiumContract, err := votiumMerkle.NewVotiumMerkle(vLog.Address, client)
			if err != nil {
				fmt.Println(err)
				continue
			}

			event, err := votiumContract.ParseClaimed(vLog)
			if err != nil {
				fmt.Println(err)
				continue
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, event.Token, event.Token, vLog.BlockNumber, 0, event.Amount, vLog.TxHash, "ethereum", "vlCVX")
		}
	}

	return bountiesClaimed
}
func fetchVotemarketV1(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config) []interfaces.BountyClaimed {
	from := config.LastBlock
	if from == 0 {
		from = 16376671
	}

	const maxBlockRange = 499
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)

	for start := from + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := start + maxBlockRange
		if end > currentBlock {
			end = currentBlock
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: VOTEMARKET_ADDRESSES_V1,
			Topics:    [][]common.Hash{{common.HexToHash("0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println("fetchVotemarketV1 logs error:", err)
			continue
		}

		for _, vLog := range logs {
			votemarketContract, err := votemarketV1.NewVotemarketV1(vLog.Address, client)
			if err != nil {
				fmt.Println("fetchVotemarketV1 contract init error:", err)
				continue
			}

			event, err := votemarketContract.ParseClaimed(vLog)
			if err != nil {
				fmt.Println("fetchVotemarketV1 parse error:", err)
				continue
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, event.RewardToken, event.RewardToken, vLog.BlockNumber, 0, event.Amount, vLog.TxHash, "ethereum", "")
		}
	}

	return bountiesClaimed
}
func fetchVotemarketV2(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config) []interfaces.BountyClaimed {
	from := config.LastBlock
	if from == 0 {
		from = 16376671
	}

	const maxBlockRange = 499
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)

	for start := from + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := start + maxBlockRange
		if end > currentBlock {
			end = currentBlock
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: VOTEMARKET_ADDRESSES_V2,
			Topics:    [][]common.Hash{{common.HexToHash("0x6f9c9826be5976f3f82a3490c52a83328ce2ec7be9e62dcb39c26da5148d7c76")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println("fetchVotemarketV2 logs error:", err)
			continue
		}

		for _, vLog := range logs {
			votemarketContract, err := votemarketV2.NewVotemarketV2(vLog.Address, client)
			if err != nil {
				fmt.Println("fetchVotemarketV2 contract init error:", err)
				continue
			}

			event, err := votemarketContract.ParseClaimed(vLog)
			if err != nil {
				fmt.Println("fetchVotemarketV2 parse error:", err)
				continue
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, event.RewardToken, event.RewardToken, vLog.BlockNumber, 0, event.Amount, vLog.TxHash, "ethereum", "")
		}
	}

	return bountiesClaimed
}

func fetchVotemarketVmV2V1(curvePools []interfaces.CurvePool, drpcKey string) []interfaces.BountyClaimed {
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)
	const maxBlockRange = 499

	for _, vmv2Config := range VOTEMARKET_VMV2_ADDRESSES_V1 {
		log.Println("Fetching bounties vm v2 on chain ", vmv2Config.ChainId)

		client, err := ethclient.Dial("https://lb.drpc.org/ogrpc?network=" + vmv2Config.DrpcChainName + "&dkey=" + drpcKey)
		if err != nil {
			log.Println(err)
			continue
		}

		currentBlockNumber, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Println("currentBlockNumber", err)
			continue
		}

		for _, market := range vmv2Config.Markets {
			for _, marketAddress := range market.Addresses {
				configFilePath := "./data/configs/bounties-config-vm-v2-" + strconv.Itoa(vmv2Config.ChainId) + "-" + marketAddress.Hex() + ".json"
				config := utils.ReadConfig(configFilePath)
				utils.WriteConfig(config, currentBlockNumber, configFilePath)

				from := config.LastBlock
				if from == 0 {
					from = uint64(market.BlockNumberDeployed)
				}

				log.Println(from, currentBlockNumber)

				for start := from + 1; start <= currentBlockNumber; start += maxBlockRange + 1 {
					end := start + maxBlockRange
					if end > currentBlockNumber {
						end = currentBlockNumber
					}

					query := ethereum.FilterQuery{
						FromBlock: big.NewInt(int64(start)),
						ToBlock:   big.NewInt(int64(end)),
						Addresses: []common.Address{marketAddress},
						Topics:    [][]common.Hash{{common.HexToHash("0x318e0a24a7fc05b12e358902d9d58475434a01768f87a4319fb35dc5b533e986")}},
					}

					logs, err := client.FilterLogs(context.Background(), query)
					if err != nil {
						log.Println("FilterLogs", err)
						continue
					}

					votemarketContract, err := vmV2.NewVmV2(marketAddress, client)
					if err != nil {
						log.Println(err)
						continue
					}

					remoteManagerAddress, err := votemarketContract.Remote(nil)
					if err != nil {
						log.Println(err)
						continue
					}

					remoteContract, err := vmRemoteManager.NewVmRemoteManager(remoteManagerAddress, client)
					if err != nil {
						log.Println(err)
						continue
					}

					tokenFactoryAddress, err := remoteContract.TOKENFACTORY(nil)
					if err != nil {
						log.Println(err)
						continue
					}

					tokenFactoryContract, err := vmTokenFactory.NewVmTokenFactory(tokenFactoryAddress, client)
					if err != nil {
						log.Println(err)
						continue
					}

					for _, vLog := range logs {
						event, err := votemarketContract.ParseClaim(vLog)
						if err != nil {
							log.Println("ParseClaim", err)
							continue
						}

						campaign, err := votemarketContract.CampaignById(nil, event.CampaignId)
						if err != nil {
							log.Println("CampaignById", err)
							continue
						}

						nativeToken, err := tokenFactoryContract.NativeTokens(nil, campaign.RewardToken)
						if err != nil {
							log.Println("NativeTokens", err)
							continue
						}

						var realTokenReward common.Address
						tokenRewardChain := ""

						if utils.IsNullAddress(nativeToken) {
							tokenRewardChain = vmv2Config.DefilammaChainName
							realTokenReward = campaign.RewardToken
						} else {
							tokenRewardChain = "ethereum"
							realTokenReward = nativeToken
						}

						if len(tokenRewardChain) == 0 {
							continue
						}

						bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, realTokenReward, campaign.RewardToken, vLog.BlockNumber, event.Epoch.Uint64(), event.Amount, vLog.TxHash, tokenRewardChain, "")
					}
				}
			}
		}
	}

	log.Println("Fetched", len(bountiesClaimed), "claimed on vm v2")
	return bountiesClaimed
}

func fetchQuest(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config) []interfaces.BountyClaimed {
	from := config.LastBlock
	if from == 0 {
		from = 14784920
	}

	const maxBlockRange = 499
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)

	for start := from + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := min(start+maxBlockRange, currentBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: QUEST_ADDRESSES,
			Topics:    [][]common.Hash{{common.HexToHash("0x9a5376f7dcf8631c2b6249c9bec3d715cb97bdd4c82d92e55d147f6b4eea4197")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println("fetchQuest logs error:", err)
			continue
		}

		for _, vLog := range logs {
			questContract, err := questDistributor.NewQuestDistributor(vLog.Address, client)
			if err != nil {
				fmt.Println("fetchQuest contract init error:", err)
				continue
			}

			event, err := questContract.ParseClaimed(vLog)
			if err != nil {
				fmt.Println("fetchQuest parse error:", err)
				continue
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, event.RewardToken, event.RewardToken, vLog.BlockNumber, 0, event.Amount, vLog.TxHash, "ethereum", "")
		}
	}

	return bountiesClaimed
}
func fetchYBribe(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config, alchemyRpcUrl string) []interfaces.BountyClaimed {
	client2, err := ethclient.Dial(alchemyRpcUrl)
	if err != nil {
		panic(err)
	}

	from := config.LastBlock
	if from == 0 {
		from = 15878261
	}

	const maxBlockRange = 499
	bountiesClaimed := make([]interfaces.BountyClaimed, 0)

	for start := from + 1; start <= currentBlock; start += maxBlockRange + 1 {
		end := min(start+maxBlockRange, currentBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: YBRIBE_ADDRESSES,
			Topics:    [][]common.Hash{{common.HexToHash("0x2422cac5e23c46c890fdcf42d0c64757409df6832174df639337558f09d99c68")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			fmt.Println("fetchYBribe logs error:", err)
			continue
		}

		for _, vLog := range logs {
			ybribeContract, err := yBribeV3.NewYBribeV3(vLog.Address, client)
			if err != nil {
				fmt.Println("yBribe contract init error:", err)
				continue
			}

			event, err := ybribeContract.ParseRewardClaimed(vLog)
			if err != nil {
				fmt.Println("parse RewardClaimed error:", err)
				continue
			}

			receipt, err := client2.TransactionReceipt(context.Background(), vLog.TxHash)
			if err != nil {
				fmt.Println("receipt fetch error:", err)
				continue
			}

			found := false
			for _, receiptLog := range receipt.Logs {
				if strings.EqualFold(receiptLog.Topics[0].Hex(), "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef") {
					t, err := erc20.NewErc20(receiptLog.Address, client)
					if err != nil {
						continue
					}

					transfert, err := t.ParseTransfer(*receiptLog)
					if err != nil {
						continue
					}

					if transfert.Value.Cmp(event.Amount) != 0 {
						continue
					}

					event.RewardToken = receiptLog.Address
					found = true
					break
				}
			}

			if !found {
				continue
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, vLog.Address, event.RewardToken, event.RewardToken, vLog.BlockNumber, 0, event.Amount, vLog.TxHash, "ethereum", "")
		}
	}

	return bountiesClaimed
}

func fetchBribeCrvFinance(client *ethclient.Client, curvePools []interfaces.CurvePool, currentBlock uint64, config interfaces.Config, alchemyRpcUrl string) []interfaces.BountyClaimed {

	bountiesClaimed := make([]interfaces.BountyClaimed, 0)
	pageKey := ""

	hexLastBlock := fmt.Sprintf("%x", config.LastBlock)
	hexCurrentBlock := fmt.Sprintf("%x", currentBlock)

	for {

		query := `{
			"id": 1,
			"jsonrpc": "2.0",
			"method": "alchemy_getAssetTransfers",
			"params": [
			{
				"fromBlock": "0x` + hexLastBlock + `",
				"toBlock": "0x` + hexCurrentBlock + `",
				"withMetadata": false,
				"excludeZeroValue": true,
				"maxCount": "0x3e8",
				"fromAddress": "0x7893bbb46613d7a4FbcC31Dab4C9b823FfeE1026",
				"category": [
				"erc20"
				],
				"order": "asc"
		`
		if len(pageKey) > 0 {
			query += ",\"pageKey\":\"" + pageKey + "\""
		}

		query += `
				}
				]
			}
		`

		body := []byte(query)

		r, err := http.NewRequest("POST", alchemyRpcUrl, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("accept", "application/json")

		httpClient := &http.Client{}
		res, err := httpClient.Do(r)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		post := new(interfaces.AlchemyTransfers)
		derr := json.NewDecoder(res.Body).Decode(post)
		if derr != nil {
			panic(derr)
		}

		if res.StatusCode != http.StatusOK {
			panic(res.Status)
		}

		pageKey = post.Result.PageKey

		for _, transfer := range post.Result.Transfers {
			blockNumber, err := strconv.ParseInt(transfer.BlockNum[2:], 16, 64)
			if err != nil {
				panic(err)
			}

			value, success := big.NewInt(0).SetString(transfer.RawContract.Value[2:], 16)
			if !success {
				panic("Error convert hex")
			}

			bountiesClaimed = addClaim(client, curvePools, bountiesClaimed, common.HexToAddress(transfer.From), common.HexToAddress(transfer.RawContract.Address), common.HexToAddress(transfer.RawContract.Address), uint64(blockNumber), 0, value, common.HexToHash(transfer.Hash), "ethereum", "")
		}

		if len(pageKey) == 0 {
			break
		}
	}

	return bountiesClaimed
}

func addClaim(client *ethclient.Client, curvePools []interfaces.CurvePool, bountiesClaimed []interfaces.BountyClaimed, contract common.Address, rewardToken common.Address, rewardTokenForDecimals common.Address, blockNumber uint64, blockTimestamp uint64, amount *big.Int, txHash common.Hash, tokenRewardChain string, comment string) []interfaces.BountyClaimed {
	decimals, err := utils.GetTokenDecimals(client, tokenRewardChain, rewardTokenForDecimals)
	if err != nil {
		fmt.Println(err)
		return bountiesClaimed
	}

	// On L2, we can't use #BlockByNumber
	// Try to fetch it and if we got an error, use the default timestamp in parameter
	// If the defaut param is set to 0, panic
	timestamp := uint64(0)
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		if blockTimestamp > 0 {
			timestamp = blockTimestamp
		} else {
			panic(err)
		}
	} else {
		timestamp = block.Time()
	}

	price := utils.GetHistoricalPriceTokenPrice(rewardToken, tokenRewardChain, timestamp)
	if price == 0 {
		price = utils.GetPriceFromCurvePools(rewardToken, tokenRewardChain, curvePools)
	}

	bountiesClaimed = append(bountiesClaimed, interfaces.BountyClaimed{
		ProtocolName: getNameBreakdown(contract),
		Timestamp:    timestamp,
		AmountUSD:    price * utils.Quo(amount, uint64(decimals)),
		Comment:      comment,
	})

	return bountiesClaimed
}

func addVotiumClaim(client *ethclient.Client, bountiesClaimed []interfaces.BountyClaimed, contract common.Address, rewardToken common.Address, blockNumber uint64, amount *big.Int, txHash common.Hash, comment string) []interfaces.BountyClaimed {
	decimals, err := utils.GetTokenDecimals(client, "ethereum", rewardToken)
	if err != nil {
		fmt.Println(err)
		return bountiesClaimed
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	if err != nil {
		panic(err)
	}

	timestamp := block.Time()
	historicalPrice := utils.GetHistoricalPriceTokenPrice(rewardToken, "ethereum", timestamp)

	nextTimestamp := timestamp + utils.WEEK_TO_SEC
	now := uint64(time.Now().Unix())
	if nextTimestamp > now {
		nextTimestamp = now
	}
	nextHistoricalPrice := utils.GetHistoricalPriceTokenPrice(rewardToken, "ethereum", nextTimestamp)

	realAmount := new(big.Int).Quo(amount, big.NewInt(2))

	bountiesClaimed = append(bountiesClaimed, interfaces.BountyClaimed{
		ProtocolName: getNameBreakdown(contract),
		Timestamp:    timestamp,
		AmountUSD:    historicalPrice * utils.Quo(realAmount, uint64(decimals)),
		Comment:      comment,
	})

	bountiesClaimed = append(bountiesClaimed, interfaces.BountyClaimed{
		ProtocolName: getNameBreakdown(contract),
		Timestamp:    nextTimestamp,
		AmountUSD:    nextHistoricalPrice * utils.Quo(realAmount, uint64(decimals)),
		Comment:      comment,
	})

	return bountiesClaimed
}

func generateDaysData(allClaimed []interfaces.BountyClaimed, start uint64, interval uint64) interfaces.StatsClaim {
	var statsClaim interfaces.StatsClaim
	statsClaim.Claims = make([]interfaces.Claim, 0)

	current := start
	currentTotalBountyDollarAmount := 0.0
	breakdown := make([]interfaces.StatsClaimBreakdown, 0)

	for _, claimed := range allClaimed {

		bountyDollarAmount := claimed.AmountUSD
		currentTotalBountyDollarAmount += bountyDollarAmount

		if claimed.Timestamp < start {
			continue
		}

		statsClaim.TotalClaimed += bountyDollarAmount
		vlCVX := isVlCVX(claimed)
		if vlCVX {
			statsClaim.VlCVXTotalClaimed += bountyDollarAmount
		} else {
			statsClaim.VeCRVTotalClaimed += bountyDollarAmount
		}

		if current+interval < claimed.Timestamp {
			statsClaim.Claims = append(statsClaim.Claims, interfaces.Claim{
				Timestamp: current,
				Value:     currentTotalBountyDollarAmount,
			})

			current += interval
		}

		if !vlCVX {
			name := claimed.ProtocolName

			found := false
			for i := 0; i < len(breakdown); i++ {
				if strings.EqualFold(breakdown[i].Name, name) {
					breakdown[i].ClaimedUSD += bountyDollarAmount
					found = true
					break
				}
			}

			if !found {
				breakdown = append(breakdown, interfaces.StatsClaimBreakdown{
					Name:       name,
					ClaimedUSD: bountyDollarAmount,
				})
			}
		}
	}

	if currentTotalBountyDollarAmount > 0 {
		statsClaim.Claims = append(statsClaim.Claims, interfaces.Claim{
			Timestamp: current,
			Value:     currentTotalBountyDollarAmount,
		})
	}

	statsClaim.VeStatsClaimBreakdown = breakdown

	return statsClaim
}

func getNameBreakdown(claimedContract common.Address) string {
	for _, contract := range VOTIUM_VE_CRV_ADDRESSES {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "Votium"
		}
	}

	for _, contract := range VOTEMARKET_ADDRESSES_V1 {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "Votemarket"
		}
	}

	for _, vm := range VOTEMARKET_VMV2_ADDRESSES_V1 {
		for _, market := range vm.Markets {
			for _, marketAddress := range market.Addresses {
				if strings.EqualFold(claimedContract.Hex(), marketAddress.Hex()) {
					return "Votemarket"
				}
			}
		}
	}

	for _, contract := range VOTEMARKET_ADDRESSES_V2 {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "Votemarket"
		}
	}

	for _, contract := range QUEST_ADDRESSES {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "Quest"
		}
	}

	for _, contract := range YBRIBE_ADDRESSES {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "YBribe"
		}
	}

	for _, contract := range BRIBE_CRV_FINANCE_ADDRESSES {
		if strings.EqualFold(claimedContract.Hex(), contract.Hex()) {
			return "bribe.crv.finance"
		}
	}

	return ""
}

func isVlCVX(bounty interfaces.BountyClaimed) bool {
	return bounty.Comment == "vlCVX"
}

func writeDataPath(claimed []interfaces.BountyClaimed) {
	file, err := json.Marshal(claimed)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(data_path, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Google Cloud
	utils.WriteBucketFile(BUCKET_BOUNTIES_DATA_FILE, claimed)
}

func readDataPath() []interfaces.BountyClaimed {
	allClaimed := make([]interfaces.BountyClaimed, 0)

	b, err := utils.ReadBucketFile(BUCKET_BOUNTIES_DATA_FILE)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &allClaimed); err != nil {
			log.Fatal(err)
		}
		return allClaimed
	}

	if !utils.FileExists(data_path) {
		return make([]interfaces.BountyClaimed, 0)
	}

	file, err := os.ReadFile(data_path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &allClaimed); err != nil {
		log.Fatal(err)
	}

	return allClaimed
}

func writeStats(stats interfaces.Stats) {
	file, err := json.Marshal(stats)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(stats_path, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Google Cloud
	utils.WriteBucketFile(BUCKET_BOUNTIES_STATS_FILE, stats)
}
