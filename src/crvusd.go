package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/crvUsdAmm"
	"main/contracts/crvUsdController"
	"main/contracts/crvUsdControllerFactory"
	"main/contracts/crvUsdMonetaryPolicy"
	"main/contracts/crvUsdOracle"
	"main/contracts/erc20"
	"main/interfaces"
	"main/utils"
	"math"
	"math/big"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var CRVUSD_CONTROLLER_FACTORY = common.HexToAddress("0xC9332fdCB1C491Dcc683bAe86Fe3cb70360738BC")

const (
	crvusdConfig          = "./data/configs/crvusd-config.json"
	crvusdControllers     = "./data/crvusd/controllers.json"
	crvusdControllersData = "./data/crvusd/controllers_data.json"
	crvusdHistoricalData  = "./data/crvusd/crvusd-historical.json"

	bucket_crvusd_dir                   = "data/crvusd"
	bucket_crvusd_controllers_file      = bucket_crvusd_dir + "/controllers.json"
	bucket_crvusd_controllers_data_file = bucket_crvusd_dir + "/controllers_data.json"
	bucket_crvusd_historical_data_file  = bucket_crvusd_dir + "/crvusd-historical.json"

	factoryBlockDeployed = 17257955
	PAS                  = 0
)

func CrvUSD(client *ethclient.Client, currentBlock uint64) {

	config := utils.ReadConfig(crvusdConfig)
	controllers := readCrvUsdControllers()
	controllersData := readCrvUsdControllerData()

	// Check if new controllers
	newControllers, err := getNewControllers(client, controllers)
	if err != nil {
		log.Fatal(err)
	}

	controllers = append(controllers, newControllers...)
	fmt.Printf("Found %d controllers\n", len(controllers))

	// Collect data
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(currentBlock)))
	if err != nil {
		panic(err)
	}

	// Create a chan to capture all collected data
	dataChan := make(chan interfaces.CrvUsdData)
	var wgDataChan sync.WaitGroup
	wgDataChan.Add(1)

	go func() {
		defer wgDataChan.Done()
		for data := range dataChan {
			fmt.Printf("Collect data received from %s at block %d\n", data.ControllerAddress.Hex(), data.BlockNumber)
			controllersData = append(controllersData, data)
		}
	}()

	// Collect data
	var wgCollect sync.WaitGroup
	for _, controller := range controllers {
		fmt.Println("Process ", controller.Address)
		wgCollect.Add(1)
		collectControllerData(&wgCollect, dataChan, client, controller, config, block)
	}

	wgCollect.Wait()
	// Close the chan, which will kill the goroutine
	close(dataChan)

	// Wait the capture go routine end
	wgDataChan.Wait()

	// Compute data
	computeData(controllersData, controllers)

	// Write controllers
	writeCrvUsdControllers(controllers)
	writeCrvUsdControllersData(controllersData)
	writeCrvUSDHistorical(controllersData, client)

	// Write config
	utils.WriteConfig(config, currentBlock, crvusdConfig)
}

func getNewControllers(client *ethclient.Client, currentControllers []interfaces.CrvUsdController) ([]interfaces.CrvUsdController, error) {
	crvUsdControllerContract, err := crvUsdControllerFactory.NewCrvUsdControllerFactory(CRVUSD_CONTROLLER_FACTORY, client)
	if err != nil {
		return nil, err
	}

	nbControllers, err := crvUsdControllerContract.NCollaterals(nil)
	if err != nil {
		return nil, err
	}

	if int(nbControllers.Int64()) == len(currentControllers) {
		return nil, nil
	}

	newControllers := make([]interfaces.CrvUsdController, 0)
	for i := int64(0); i < nbControllers.Int64(); i++ {
		controllerAddress, err := crvUsdControllerContract.Controllers(nil, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		found := false
		for _, c := range currentControllers {
			if strings.EqualFold(c.Address.Hex(), controllerAddress.Hex()) {
				found = true
				break
			}
		}

		if !found {

			controller, err := createCrvUsdController(client, controllerAddress)
			if err != nil {
				return nil, err
			}
			newControllers = append(newControllers, controller)
		}
	}
	return newControllers, nil
}

func createCrvUsdController(client *ethclient.Client, controllerAddress common.Address) (interfaces.CrvUsdController, error) {
	controllerContract, err := crvUsdController.NewCrvUsdController(controllerAddress, client)
	if err != nil {
		return interfaces.CrvUsdController{}, err
	}

	collateralTokenAddress, err := controllerContract.CollateralToken(nil)
	if err != nil {
		return interfaces.CrvUsdController{}, err
	}

	collateralToken, err := erc20.NewErc20(collateralTokenAddress, client)
	if err != nil {
		return interfaces.CrvUsdController{}, err
	}

	collateralDecimals, err := collateralToken.Decimals(nil)
	if err != nil {
		return interfaces.CrvUsdController{}, err
	}

	collateralSymbol, err := collateralToken.Symbol(nil)
	if err != nil {
		return interfaces.CrvUsdController{}, err
	}

	return interfaces.CrvUsdController{
		Address:                 controllerAddress,
		Name:                    collateralSymbol,
		CollateralTokenAddress:  collateralTokenAddress,
		CollateralTokenDecimals: collateralDecimals,
	}, nil
}

func collectControllerData(wgCollect *sync.WaitGroup, dataChan chan interfaces.CrvUsdData, client *ethclient.Client, controller interfaces.CrvUsdController, config interfaces.Config, block *types.Block) {

	defer wgCollect.Done()

	from := config.LastBlock
	if from == 0 {
		from = factoryBlockDeployed
	}

	to := block.NumberU64()

	// No historical, from workflow
	if PAS == 0 {
		crvUsdData, err := internalCollectControllerData(client, controller, to, from)

		if err != nil {
			fmt.Println(err)
		} else {
			dataChan <- *crvUsdData
		}
	} else {
		for ok := true; ok; ok = from < to {
			next := from + PAS

			fmt.Printf("%s -> %d - %d\n", controller.Address.Hex(), from, next)

			crvUsdData, err := internalCollectControllerData(client, controller, from, next)

			// Err can be contract not found
			if err != nil {
				fmt.Println(err)
			} else {
				dataChan <- *crvUsdData
			}
			from = next
		}
	}
}

func internalCollectControllerData(client *ethclient.Client, controller interfaces.CrvUsdController, from uint64, to uint64) (*interfaces.CrvUsdData, error) {
	opts := new(bind.CallOpts)
	opts.BlockNumber = big.NewInt(int64(from))

	currentBlock, err := client.BlockByNumber(context.Background(), opts.BlockNumber)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

	controllerContract, err := crvUsdController.NewCrvUsdController(controller.Address, client)
	if err != nil {
		fmt.Println("2")
		return nil, err
	}

	totalDebt, err := controllerContract.TotalDebt(opts)
	if err != nil {
		fmt.Println("3", controller.Address.Hex(), from)
		return nil, err
	}

	ammAddress, err := controllerContract.Amm(opts)
	if err != nil {
		return nil, err
	}

	monetaryAddress, err := controllerContract.MonetaryPolicy(opts)
	if err != nil {
		return nil, err
	}

	ammContract, err := crvUsdAmm.NewCrvUsdAmm(ammAddress, client)
	if err != nil {
		fmt.Println("5")
		return nil, err
	}

	controllerRate, err := ammContract.Rate(opts)
	if err != nil {
		return nil, err
	}

	adminFeesY, err := ammContract.AdminFeesY(opts)
	if err != nil {
		return nil, err
	}

	monetaryContract, err := crvUsdMonetaryPolicy.NewCrvUsdMonetaryPolicy(monetaryAddress, client)
	if err != nil {
		return nil, err
	}

	controllerFuturRate, err := monetaryContract.Rate0(opts, controller.Address)
	if err != nil {
		controllerFuturRate, err = monetaryContract.Rate00(opts)
		if err != nil {
			return nil, err
		}
	}

	collateralTokenContract, err := erc20.NewErc20(controller.CollateralTokenAddress, client)
	if err != nil {
		return nil, err
	}

	collateralTokenBalance, err := collateralTokenContract.BalanceOf(opts, ammAddress)
	if err != nil {
		return nil, err
	}

	crvUSDContract, err := erc20.NewErc20(utils.CRVUSD_ADDRESS, client)
	if err != nil {
		return nil, err
	}

	crvUSDAMMBalance, err := crvUSDContract.BalanceOf(opts, ammAddress)
	if err != nil {
		return nil, err
	}

	collateralTokenPrice, err := ammContract.PriceOracle(opts)
	if err != nil {
		return nil, err
	}

	crvUSDOracleAddress, err := monetaryContract.PRICEORACLE(opts)
	if err != nil {
		return nil, err
	}

	crvUsdOracleContract, err := crvUsdOracle.NewCrvUsdOracle(crvUSDOracleAddress, client)
	if err != nil {
		return nil, err
	}

	crvUSDPrice, err := crvUsdOracleContract.Price(opts)
	if err != nil {
		fmt.Println("13")
		return nil, err
	}

	feesPending, err := controllerContract.AdminFees(opts)
	if err != nil {
		fmt.Println("14")
		return nil, err
	}

	// Collateral value
	crvUSDPriceF := utils.Quo(crvUSDPrice, uint64(18))
	crvUSDAMMBalanceF := utils.Quo(crvUSDAMMBalance, 18)

	collateralTokenPriceF := utils.Quo(collateralTokenPrice, uint64(18))
	collateralTokenBalanceF := utils.Quo(collateralTokenBalance, uint64(controller.CollateralTokenDecimals))
	adminFeesYF := utils.Quo(adminFeesY, uint64(18))

	// Rates
	controllerRateF := utils.Quo(controllerRate, uint64(18))
	controllerFuturRateF := utils.Quo(controllerFuturRate, uint64(18))

	currentRate := (math.Pow(2.718281828459, controllerRateF*365*86400) - 1) * 100
	futureRate := (math.Pow(2.718281828459, controllerFuturRateF*365*86400) - 1) * 100

	crvUsdData := new(interfaces.CrvUsdData)
	crvUsdData.ControllerAddress = controller.Address
	crvUsdData.BlockNumber = currentBlock.NumberU64()
	crvUsdData.BlockTimestamp = currentBlock.Time()
	crvUsdData.CollateralUSD = (crvUSDAMMBalanceF * crvUSDPriceF) + (collateralTokenPriceF * collateralTokenBalanceF) + (adminFeesYF * crvUSDPriceF)
	crvUsdData.CurrentAPR = currentRate
	crvUsdData.FutureAPR = futureRate
	crvUsdData.TotalDept = utils.Quo(totalDebt, uint64(18))
	crvUsdData.FeesPendingUSD = utils.Quo(feesPending, uint64(18))
	crvUsdData.FeesCollectedUSD = getFeesCollected(client, controller, int64(from), int64(to))
	crvUsdData.AmmAddress = ammAddress
	crvUsdData.MonetaryAddress = monetaryAddress
	crvUsdData.CrvUSDOracleAddress = crvUSDOracleAddress

	return crvUsdData, nil
}
func getFeesCollected(client *ethclient.Client, controller interfaces.CrvUsdController, from int64, toBlock int64) float64 {
	if toBlock < from {
		from, toBlock = toBlock, from
	}

	var feesCollected float64

	err := utils.ForEachBlockRange(uint64(from), uint64(toBlock), 499, func(start uint64, end uint64) error {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: []common.Address{controller.Address},
			Topics:    [][]common.Hash{{common.HexToHash("0x5393ab6ef9bb40d91d1b04bbbeb707fbf3d1eb73f46744e2d179e4996026283f")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			return fmt.Errorf("getFeesCollected logs error: %w", err)
		}

		for _, vLog := range logs {
			controllerContract, err := crvUsdController.NewCrvUsdController(vLog.Address, client)
			if err != nil {
				fmt.Println("controller init error:", err)
				continue
			}

			event, err := controllerContract.ParseCollectFees(vLog)
			if err != nil {
				fmt.Println("parse CollectFees error:", err)
				continue
			}

			feesCollected += utils.Quo(event.Amount, 18)
		}
		return nil
	})

	if err != nil {
		fmt.Println("getFeesCollected error:", err)
	}

	return feesCollected
}

func computeData(controllersData []interfaces.CrvUsdData, controllers []interfaces.CrvUsdController) {
	sort.Slice(controllersData, func(i, j int) bool {
		return controllersData[i].BlockTimestamp < controllersData[j].BlockTimestamp
	})

	for i := range controllers {
		controllers[i].TotalFeesCollectedUSD = 0
	}

	for _, crvUsdData := range controllersData {
		for i := range controllers {
			if strings.EqualFold(controllers[i].Address.Hex(), crvUsdData.ControllerAddress.Hex()) {
				controllers[i].TotalFeesCollectedUSD += crvUsdData.FeesCollectedUSD
				controllers[i].CrvUsdData = crvUsdData
				break
			}
		}
	}
}

func writeCrvUSDHistorical(controllersData []interfaces.CrvUsdData, client *ethclient.Client) {
	// Total supply per day
	// Based on the last data of each day

	// Sort all ASC
	sort.Slice(controllersData, func(i, j int) bool {
		return controllersData[i].BlockTimestamp < controllersData[j].BlockTimestamp
	})

	// Split between controllers
	dataPerController := make(map[common.Address][]interfaces.CrvUsdData)
	for _, data := range controllersData {
		_, exists := dataPerController[data.ControllerAddress]
		if !exists {
			dataPerController[data.ControllerAddress] = make([]interfaces.CrvUsdData, 0)
		}

		dataPerController[data.ControllerAddress] = append(dataPerController[data.ControllerAddress], data)
	}

	// Iterate other the map
	// For each controller, at the end of each day, we will get the totalSupply to compute it
	totalSupplyPerDay := make(map[int]interfaces.CrvUSDHistoricalData)
	for _, datas := range dataPerController {
		previousTimestamp := int64(datas[0].BlockTimestamp)

		for i, data := range datas {
			if utils.IsDifferentDay(previousTimestamp, int64(data.BlockTimestamp)) {
				// Get the previous one to get the "last" day data
				previousData := datas[i-1]
				day := utils.DaySinceEpoch(int64(previousData.BlockTimestamp))

				s, exists := totalSupplyPerDay[day]
				if !exists {
					s = interfaces.CrvUSDHistoricalData{
						TotalSupply: 0,
						Timestamp:   previousData.BlockTimestamp,
						Tvl:         0,
					}
				}

				s.TotalSupply += previousData.TotalDept
				s.Tvl += previousData.CollateralUSD
				totalSupplyPerDay[day] = s

				previousTimestamp = int64(data.BlockTimestamp)
			}
		}
	}

	// Convert to struct
	historicalData := make([]interfaces.CrvUSDHistoricalData, 0)
	for _, s := range totalSupplyPerDay {
		historicalData = append(historicalData, interfaces.CrvUSDHistoricalData{
			TotalSupply: s.TotalSupply,
			Timestamp:   s.Timestamp,
			Tvl:         s.Tvl,
		})
	}

	// Sort ASC
	sort.Slice(historicalData, func(i, j int) bool {
		return historicalData[i].Timestamp < historicalData[j].Timestamp
	})

	crvUSDOracleContract, err := crvUsdOracle.NewCrvUsdOracle(common.HexToAddress("0x18672b1b0c623a30089A280Ed9256379fb0E4E62"), client)
	crvUSDPrice := 0.0
	if err == nil {
		currentCrvUSDPrice, err := crvUSDOracleContract.LastPrice(nil)
		if err == nil {
			crvUSDPrice = utils.Quo(currentCrvUSDPrice, 18)
		}
	}

	data := interfaces.CrvUSDHistorical{
		CurrentCrvUSDPrice:  crvUSDPrice,
		DailyHistoricalData: historicalData,
	}
	file, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(crvusdHistoricalData, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Google Cloud
	if err := utils.WriteBucketFile(bucket_crvusd_historical_data_file, data); err != nil {
		fmt.Println(err)
	}
}

func readCrvUsdControllerData() []interfaces.CrvUsdData {
	controllersData := make([]interfaces.CrvUsdData, 0)
	b, err := utils.ReadBucketFile(bucket_crvusd_controllers_data_file)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &controllersData); err != nil {
			log.Fatal(err)
		}
		return controllersData
	}

	if !utils.FileExists(crvusdControllersData) {
		return make([]interfaces.CrvUsdData, 0)
	}

	file, err := os.ReadFile(crvusdControllersData)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &controllersData); err != nil {
		log.Fatal(err)
	}

	return controllersData
}

func writeCrvUsdControllersData(controllersData []interfaces.CrvUsdData) {
	file, err := json.Marshal(controllersData)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(crvusdControllersData, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Google Cloud
	if err := utils.WriteBucketFile(bucket_crvusd_controllers_data_file, controllersData); err != nil {
		fmt.Println(err)
	}
}

func readCrvUsdControllers() []interfaces.CrvUsdController {

	controllers := make([]interfaces.CrvUsdController, 0)
	b, err := utils.ReadBucketFile(bucket_crvusd_controllers_file)
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &controllers); err != nil {
			log.Fatal(err)
		}
		return controllers
	}

	if !utils.FileExists(crvusdControllers) {
		return make([]interfaces.CrvUsdController, 0)
	}

	file, err := os.ReadFile(crvusdControllers)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &controllers); err != nil {
		log.Fatal(err)
	}

	return controllers
}

func writeCrvUsdControllers(controllers []interfaces.CrvUsdController) {
	file, err := json.Marshal(controllers)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(crvusdControllers, file, 0644); err != nil {
		log.Fatal(err)
	}

	// Google Cloud
	if err := utils.WriteBucketFile(bucket_crvusd_controllers_file, controllers); err != nil {
		fmt.Println(err)
	}
}
