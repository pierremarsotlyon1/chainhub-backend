package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/erc20"
	"main/contracts/llamalendController"
	"main/contracts/llamalendFactory"
	"main/contracts/llamalendVault"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	LLAMALEND_PER_DAY_PATH    = "./data/llamalend/llamalend-per-day.json"
	LLAMALEND_CONFIGS_MAINNET = "./data/configs/llamalend-mainnet.json"
	LLAMALEND_CONFIGS_ARB     = "./data/configs/llamalend-arb.json"
)

var LLAMALEND_FACTORIES = []interfaces.LlamalendConfig{
	{
		FactoryAddress:  common.HexToAddress("0xeA6876DDE9e3467564acBeE1Ed5bac88783205E0"),
		ChainId:         1,
		MarketName:      "ethereum",
		TimestampDeploy: 1710294146,
		BlockDeploy:     19422660,
		RpcUrl:          "https://eth-mainnet.g.alchemy.com/v2/" + utils.GoDotEnvVariable("ALCHEMY_APIKEY_3"),
		ConfigPath:      LLAMALEND_CONFIGS_MAINNET,
	},
	{
		FactoryAddress:  common.HexToAddress("0xcaEC110C784c9DF37240a8Ce096D352A75922DeA"),
		ChainId:         42161,
		MarketName:      "arbitrum",
		TimestampDeploy: 1711266146,
		BlockDeploy:     193652535,
		RpcUrl:          "https://arb-mainnet.g.alchemy.com/v2/" + utils.GoDotEnvVariable("ALCHEMY_APIKEY_3"),
		ConfigPath:      LLAMALEND_CONFIGS_ARB,
	},
}

// Caches
var tokenSymbol = make(map[common.Address]string)
var tokenDecimals = make(map[common.Address]uint8)

func Llamalend() {

	curvePools, err := utils.GetAllCurvePools()
	if err != nil {
		log.Fatal(err)
	}

	currentMarkets := readMarkets()
	controllersMapHardLiq := make(map[common.Address]bool)

	now := uint64(time.Now().Unix())

	for _, factory := range LLAMALEND_FACTORIES {
		client, err := ethclient.Dial(factory.RpcUrl)
		if err != nil {
			log.Fatal(err)
		}

		blockNumber := utils.GetBlockNumberByTimestamp(factory.MarketName, now)

		config := utils.ReadConfig(factory.ConfigPath)

		opts := new(bind.CallOpts)
		opts.BlockNumber = big.NewInt(int64(blockNumber))

		blockTimestamp := now
		dayTimestamp := utils.GetStartOfDay(uint64(blockTimestamp))

		markets := fetchLlamalend(opts, factory, client, blockTimestamp, dayTimestamp, curvePools)

		for _, market := range markets {

			found := false
			for i := 0; i < len(currentMarkets); i++ {
				if strings.EqualFold(currentMarkets[i].FactoryAddress.Hex(), market.FactoryAddress.Hex()) && strings.EqualFold(currentMarkets[i].ControllerAddress.Hex(), market.ControllerAddress.Hex()) && market.Id == currentMarkets[i].Id && market.Timestamp == currentMarkets[i].Timestamp {
					currentMarkets[i].TotalDebt = market.TotalDebt
					currentMarkets[i].TotalDebtUSD = market.TotalDebtUSD
					currentMarkets[i].Supplied = market.Supplied
					currentMarkets[i].SuppliedUSD = market.SuppliedUSD
					currentMarkets[i].Collateral = market.Collateral
					currentMarkets[i].CollateralUSD = market.CollateralUSD
					currentMarkets[i].BorrowedTokenAddress = market.BorrowedTokenAddress
					currentMarkets[i].CollateralTokenAddress = market.CollateralTokenAddress
					currentMarkets[i].ChainId = market.ChainId
					currentMarkets[i].Loans = market.Loans
					currentMarkets[i].BorrowApr = market.BorrowApr
					currentMarkets[i].SupplyApr = market.SupplyApr

					found = true
				}
			}

			if !found {
				currentMarkets = append(currentMarkets, market)
			}
		}

		for _, market := range currentMarkets {
			if market.ChainId != factory.ChainId {
				continue
			}
			_, exists := controllersMapHardLiq[market.ControllerAddress]

			blockFrom := config.LastBlock

			if !exists {
				from := int64(blockFrom)

				liquidations := fetchHardLiquidation(client, now, market.ControllerAddress, from, int64(blockNumber))

				for _, liquidation := range liquidations {
					timestamp := utils.GetStartOfDay(liquidation.Timestamp)

					for i := 0; i < len(currentMarkets); i++ {
						if strings.EqualFold(currentMarkets[i].ControllerAddress.Hex(), market.ControllerAddress.Hex()) && currentMarkets[i].Timestamp == uint64(timestamp) {
							currentMarkets[i].HardLiquidation = currentMarkets[i].HardLiquidation + 1
							break
						}
					}
				}

				// For Alchemy rate limite
				time.Sleep(2 * time.Second)
				controllersMapHardLiq[market.ControllerAddress] = true
			}
		}

		utils.WriteConfig(config, blockNumber, factory.ConfigPath)
	}

	writeMarkets(currentMarkets)
}

func LlamalendHistorical() {

	curvePools, err := utils.GetAllCurvePools()
	if err != nil {
		log.Fatal(err)
	}

	configsMainnet := utils.ReadConfig(LLAMALEND_CONFIGS_MAINNET)
	configsArb := utils.ReadConfig(LLAMALEND_CONFIGS_ARB)

	markets := make([]interfaces.LlamalendMarket, 0)

	for _, factory := range LLAMALEND_FACTORIES {
		client, err := ethclient.Dial(factory.RpcUrl)
		if err != nil {
			log.Fatal(err)
		}

		block, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}

		factoryMarkets := make([]interfaces.LlamalendMarket, 0)

		for blockTimestamp := factory.TimestampDeploy + utils.DAY_TO_SEC; blockTimestamp < block.Time; blockTimestamp += utils.DAY_TO_SEC {

			blockNumber := utils.GetBlockNumberByTimestamp(factory.MarketName, blockTimestamp)
			dayTimestamp := utils.GetStartOfDay(uint64(blockTimestamp))
			fmt.Println("dayTimestamp :", dayTimestamp)
			fmt.Println("blockNumber :", blockNumber)
			fmt.Println("-----")

			opts := new(bind.CallOpts)
			opts.BlockNumber = big.NewInt(int64(blockNumber))

			factoryMarkets = append(markets, fetchLlamalend(opts, factory, client, blockTimestamp, dayTimestamp, curvePools)...)
		}

		markets = append(markets, factoryMarkets...)

		controllersMap := make(map[common.Address]bool)
		for _, market := range factoryMarkets {
			_, exists := controllersMap[market.ControllerAddress]
			if !exists {
				liquidations := fetchHardLiquidation(client, uint64(time.Now().Unix()), market.ControllerAddress, int64(factory.BlockDeploy), block.Number.Int64())
				for _, liquidation := range liquidations {
					timestamp := utils.GetStartOfDay(liquidation.Timestamp)

					for i := 0; i < len(factoryMarkets); i++ {
						if strings.EqualFold(factoryMarkets[i].ControllerAddress.Hex(), market.ControllerAddress.Hex()) && factoryMarkets[i].Timestamp == uint64(timestamp) {
							factoryMarkets[i].HardLiquidation = factoryMarkets[i].HardLiquidation + 1
							break
						}
					}

				}
				controllersMap[market.ControllerAddress] = true
			}
		}

		// Set configs
		if strings.EqualFold(factory.FactoryAddress.Hex(), "0xeA6876DDE9e3467564acBeE1Ed5bac88783205E0") {
			// Mainnet
			utils.WriteConfig(configsMainnet, block.Number.Uint64(), LLAMALEND_CONFIGS_MAINNET)
		} else if strings.EqualFold(factory.FactoryAddress.Hex(), "0xcaEC110C784c9DF37240a8Ce096D352A75922DeA") {
			// Arb
			utils.WriteConfig(configsArb, block.Number.Uint64(), LLAMALEND_CONFIGS_ARB)
		}
	}

	writeMarkets(markets)
}

func fetchLlamalend(opts *bind.CallOpts, factory interfaces.LlamalendConfig, client *ethclient.Client, blockTimestamp uint64, dayTimestamp int64, curvePools []interfaces.CurvePool) []interfaces.LlamalendMarket {

	markets := make([]interfaces.LlamalendMarket, 0)

	factoryContract, err := llamalendFactory.NewLlamalendFactory(factory.FactoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	marketCount, err := factoryContract.MarketCount(opts)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < int(marketCount.Int64()); i++ {
		controllerAddress, err := factoryContract.Controllers(opts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		vaultAddress, err := factoryContract.Vaults(opts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		ammAddress, err := factoryContract.Amms(opts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		controllerContract, err := llamalendController.NewLlamalendController(controllerAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		vaultContract, err := llamalendVault.NewLlamalendVault(vaultAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		borrowAprBN, err := vaultContract.BorrowApr(opts)
		if err != nil {
			log.Fatal(err)
		}

		supplyBN, err := vaultContract.LendApr(opts)
		if err != nil {
			log.Fatal(err)
		}

		nbLoans, err := controllerContract.NLoans(opts)
		if err != nil {
			log.Fatal(err)
		}

		totalDebtBN, err := controllerContract.TotalDebt(opts)
		if err != nil {
			log.Fatal(err)
		}

		collateralTokenAddress, err := controllerContract.CollateralToken(opts)
		if err != nil {
			log.Fatal(err)
		}

		collateralTokenContract, err := erc20.NewErc20(collateralTokenAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		collateralTokenSymbol, exists := tokenSymbol[collateralTokenAddress]
		if !exists {

			collateralTokenSymbol, err = collateralTokenContract.Symbol(opts)
			if err != nil {
				log.Fatal(err)
			}

			tokenSymbol[collateralTokenAddress] = collateralTokenSymbol
		}

		collateralTokenDecimals, exists := tokenDecimals[collateralTokenAddress]
		if !exists {

			collateralTokenDecimals, err = collateralTokenContract.Decimals(opts)
			if err != nil {
				log.Fatal(err)
			}

			tokenDecimals[collateralTokenAddress] = collateralTokenDecimals
		}

		// Borrowed
		borrowedTokenAddress, err := controllerContract.BorrowedToken(opts)
		if err != nil {
			log.Fatal(err)
		}

		borrowTokenContract, err := erc20.NewErc20(borrowedTokenAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		borrowTokenSymbol, exists := tokenSymbol[borrowedTokenAddress]
		if !exists {

			borrowTokenSymbol, err = borrowTokenContract.Symbol(opts)
			if err != nil {
				log.Fatal(err)
			}

			tokenSymbol[borrowedTokenAddress] = borrowTokenSymbol
		}

		borrowTokenDecimals, exists := tokenDecimals[borrowedTokenAddress]
		if !exists {

			borrowTokenDecimals, err = borrowTokenContract.Decimals(opts)
			if err != nil {
				log.Fatal(err)
			}

			tokenDecimals[borrowedTokenAddress] = borrowTokenDecimals
		}

		// Collateral
		collateralBN, err := collateralTokenContract.BalanceOf(opts, ammAddress)
		if err != nil {
			log.Fatal(err)
		}

		availableToBorrowBN, err := borrowTokenContract.BalanceOf(opts, controllerAddress)
		if err != nil {
			log.Fatal(err)
		}

		borrowTokenPrice := getTokenPrice(borrowedTokenAddress, factory.MarketName, blockTimestamp, curvePools)
		collateralTokenPrice := getTokenPrice(collateralTokenAddress, factory.MarketName, blockTimestamp, curvePools)

		totalDebt := utils.Quo(totalDebtBN, uint64(borrowTokenDecimals))
		totalDebtUSD := totalDebt * borrowTokenPrice

		supplied := totalDebt + utils.Quo(availableToBorrowBN, uint64(borrowTokenDecimals))
		suppliedUSD := supplied * borrowTokenPrice

		collateral := utils.Quo(collateralBN, uint64(collateralTokenDecimals))
		collateralUSD := collateral * collateralTokenPrice

		markets = append(markets, interfaces.LlamalendMarket{
			Id:                     i,
			ControllerAddress:      controllerAddress,
			FactoryAddress:         factory.FactoryAddress,
			MarketName:             collateralTokenSymbol + "/" + borrowTokenSymbol,
			Timestamp:              uint64(dayTimestamp),
			TotalDebt:              totalDebt,
			TotalDebtUSD:           totalDebtUSD,
			Supplied:               supplied,
			SuppliedUSD:            suppliedUSD,
			Collateral:             collateral,
			CollateralUSD:          collateralUSD,
			BorrowedTokenAddress:   borrowedTokenAddress,
			CollateralTokenAddress: collateralTokenAddress,
			ChainId:                factory.ChainId,
			Loans:                  nbLoans.Int64(),
			BorrowApr:              utils.Quo(borrowAprBN, 18),
			SupplyApr:              utils.Quo(supplyBN, 18),
		})
	}

	return markets
}

func fetchHardLiquidation(client *ethclient.Client, now uint64, controllerAddress common.Address, blockFrom int64, blockTo int64) []interfaces.LlamalendLiquidate {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockFrom),
		ToBlock:   big.NewInt(blockTo),
		Addresses: []common.Address{controllerAddress},
		Topics:    [][]common.Hash{{common.HexToHash("0x642dd4d37ddd32036b9797cec464c0045dd2118c549066ae6b0f88e32240c2d0")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	llamalendLiquidates := make([]interfaces.LlamalendLiquidate, 0)

	for _, vLog := range logs {
		controllerContract, err := llamalendController.NewLlamalendController(vLog.Address, client)
		if err != nil {
			panic(err)
		}

		event, err := controllerContract.ParseLiquidate(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		llamalendLiquidates = append(llamalendLiquidates, interfaces.LlamalendLiquidate{
			Timestamp: now,
			User:      event.User,
		})
	}

	return llamalendLiquidates
}

func getTokenPrice(tokenAddress common.Address, chainName string, timestamp uint64, curvePools []interfaces.CurvePool) float64 {
	tokenPrice := utils.GetHistoricalPriceTokenPrice(tokenAddress, chainName, timestamp)
	if tokenPrice == 0 {
		for _, pool := range curvePools {
			for _, coin := range pool.Coins {
				if strings.EqualFold(coin.Address, tokenAddress.Hex()) {
					return coin.UsdPrice.(float64)
				}
			}
		}
	}

	return tokenPrice
}

func readMarkets() []interfaces.LlamalendMarket {

	if !utils.FileExists(LLAMALEND_PER_DAY_PATH) {
		return make([]interfaces.LlamalendMarket, 0)
	}

	file, err := os.ReadFile(LLAMALEND_PER_DAY_PATH)
	if err != nil {
		log.Fatal(err)
	}

	markets := make([]interfaces.LlamalendMarket, 0)
	if err := json.Unmarshal([]byte(file), &markets); err != nil {
		log.Fatal(err)
	}

	return markets
}

func writeMarkets(markets []interfaces.LlamalendMarket) {
	file, err := json.Marshal(markets)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(LLAMALEND_PER_DAY_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}
