package src

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/crvUsdController"
	"main/contracts/crvUsdControllerFactory"
	"main/contracts/crvUsdMonetaryPolicy"
	"main/contracts/erc20"
	"main/contracts/pegKeeperPool"
	pegKeeper "main/contracts/pegkeeper"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	PEGKEEPERS_CONFIG     = "./data/configs/pegkeepers-config.json"
	PEGKEEPER_DATA        = "./data/pegkeepers/data.json"
	PEGKEEPER_VOLUME_DATA = "./data/pegkeepers/volume.json"
	PEGKEEPER_FEE_DATA    = "./data/pegkeepers/fee.json"
)

func FetchPegKeepers(client *ethclient.Client, currentBlock uint64, currentBlockTImestamp uint64) {
	pegKeeperAddresses, err := getPegKeeperAddresses(client)
	if err != nil {
		log.Fatal(err)
	}

	config := utils.ReadConfig(PEGKEEPERS_CONFIG)

	opts := new(bind.CallOpts)
	opts.BlockNumber = big.NewInt(int64(currentBlock))

	// DATA
	pegKeepers := readPegKeepers()
	fees := readPegKeepersFees()

	// Run channels
	pegkeeperChan := make(chan interfaces.PegKeeper)
	feesChan := make(chan interfaces.PegKeeperFee)

	var wgDataChan sync.WaitGroup
	wgDataChan.Add(1)

	go func() {
		defer wgDataChan.Done()
		for data := range pegkeeperChan {

			found := false
			for i := 0; i < len(pegKeepers); i++ {
				if strings.EqualFold(pegKeepers[i].Address.Hex(), data.Address.Hex()) {
					pegKeepers[i].Debt = data.Debt
					pegKeepers[i].Tvl = data.Tvl
					pegKeepers[i].Name = data.Name
					pegKeepers[i].PoolAddress = data.PoolAddress
					pegKeepers[i].FeesPending = data.FeesPending
					found = true
					break
				}
			}

			if !found {
				pegKeepers = append(pegKeepers, data)
			}
		}
	}()

	var wgFeeChan sync.WaitGroup
	wgFeeChan.Add(1)

	go func() {
		defer wgFeeChan.Done()
		for data := range feesChan {
			fees = append(fees, data)
		}
	}()

	// Collect data
	var wgCollectData sync.WaitGroup
	for _, pegKeeperAddress := range pegKeeperAddresses {
		wgCollectData.Add(1)
		collectData(&wgCollectData, feesChan, pegkeeperChan, client, pegKeeperAddress, opts, config)
	}

	wgCollectData.Wait()
	close(pegkeeperChan)
	close(feesChan)

	wgDataChan.Wait()
	wgFeeChan.Wait()

	// Reset pegkeeper totals
	for i := range pegKeepers {
		pegKeepers[i].FeesCollected = 0
	}

	// Compute fees
	for _, fee := range fees {
		for i, pegKeeper := range pegKeepers {
			if strings.EqualFold(fee.PegKeeperAddress.Hex(), pegKeeper.Address.Hex()) {
				pegKeepers[i].FeesCollected += (fee.Amount * fee.LpPrice)
				break
			}
		}
	}

	writePegKeepers(pegKeepers)
	writePegKeepersFees(fees)

	// Write config
	utils.WriteConfig(config, currentBlock, PEGKEEPERS_CONFIG)
}

func getPegKeeperAddresses(client *ethclient.Client) ([]common.Address, error) {
	crvUsdControllerContract, err := crvUsdControllerFactory.NewCrvUsdControllerFactory(CRVUSD_CONTROLLER_FACTORY, client)
	if err != nil {
		return nil, err
	}

	nbControllers, err := crvUsdControllerContract.NCollaterals(nil)
	if err != nil {
		return nil, err
	}

	pegKeeperAddresses := make(map[common.Address]bool)

	for i := int64(0); i < nbControllers.Int64(); i++ {
		controllerAddress, err := crvUsdControllerContract.Controllers(nil, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		controllerContract, err := crvUsdController.NewCrvUsdController(controllerAddress, client)
		if err != nil {
			return nil, err
		}

		monetaryAddress, err := controllerContract.MonetaryPolicy(nil)
		if err != nil {
			return nil, err
		}

		monetaryContract, err := crvUsdMonetaryPolicy.NewCrvUsdMonetaryPolicy(monetaryAddress, client)
		if err != nil {
			return nil, err
		}

		indexPegKeeper := 0
		for {
			pegKeeperAddress, err := monetaryContract.PegKeepers(nil, big.NewInt(int64(indexPegKeeper)))
			if err != nil {
				return nil, err
			}

			if utils.IsNullAddress(pegKeeperAddress) {
				break
			}

			pegKeeperAddresses[pegKeeperAddress] = true
			indexPegKeeper++
		}
	}

	pegKeeperAddressesArray := make([]common.Address, 0)
	for pegKeeperAddress := range pegKeeperAddresses {
		pegKeeperAddressesArray = append(pegKeeperAddressesArray, pegKeeperAddress)
	}
	return pegKeeperAddressesArray, nil
}

func collectData(wgCollectData *sync.WaitGroup, feesChan chan interfaces.PegKeeperFee, dataChan chan interfaces.PegKeeper, client *ethclient.Client, pegKeeperAddress common.Address, opts *bind.CallOpts, config interfaces.Config) {
	defer wgCollectData.Done()
	pegKeeperContract, err := pegKeeper.NewPegKeeper(pegKeeperAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	poolAddress, err := pegKeeperContract.Pool(opts)
	if err != nil {
		log.Fatal(err)
	}

	pegKeeperPoolContract, err := pegKeeperPool.NewPegKeeperPool(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	poolName, err := pegKeeperPoolContract.Name(opts)
	if err != nil {
		log.Fatal(err)
	}

	name := poolName
	indexDot := strings.Index(poolName, ": ")
	if indexDot != -1 {
		name = strings.Trim(poolName[indexDot+2:], "")
	}

	indexCoin := int64(0)
	tvl := 0.0
	for {
		coinAddress, err := pegKeeperPoolContract.Coins(opts, big.NewInt(indexCoin))
		if err != nil {
			break
		}

		coinContract, err := erc20.NewErc20(coinAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		poolCoinBalance, err := coinContract.BalanceOf(opts, poolAddress)
		if err != nil {
			log.Fatal(err)
		}

		coinDecimals, err := coinContract.Decimals(opts)
		if err != nil {
			log.Fatal(err)
		}

		poolCoinBalanceF := utils.Quo(poolCoinBalance, uint64(coinDecimals))
		tokenPrice := utils.GetHistoricalPriceTokenPrice(coinAddress, "ethereum", uint64(time.Now().Unix()))

		tvl += (tokenPrice * poolCoinBalanceF)
		indexCoin++
	}

	debt, err := pegKeeperContract.Debt(opts)
	if err != nil {
		log.Fatal(err)
	}

	currentProfit, err := pegKeeperContract.CalcProfit(opts)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch volume
	getFees(feesChan, client, pegKeeperAddress, opts, config)

	dataChan <- interfaces.PegKeeper{
		Address:       pegKeeperAddress,
		PoolAddress:   poolAddress,
		Name:          name,
		Tvl:           tvl,
		Debt:          utils.Quo(debt, 18),
		FeesCollected: 0,
		FeesPending:   utils.Quo(currentProfit, 18),
	}
}

func getFees(feesChan chan interfaces.PegKeeperFee, client *ethclient.Client, pegKeeperAddress common.Address, opts *bind.CallOpts, config interfaces.Config) {
	from := config.LastBlock
	if from == 0 {
		from = 17258030
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from) + 1),
		ToBlock:   opts.BlockNumber,
		Addresses: []common.Address{pegKeeperAddress},
		Topics:    [][]common.Hash{{common.HexToHash("0x357d905f1831209797df4d55d79c5c5bf1d9f7311c976afd05e13d881eab9bc8")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return
	}

	pegKeeperContract, err := pegKeeper.NewPegKeeper(pegKeeperAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	poolAddress, err := pegKeeperContract.Pool(opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println("BlockByNumber", poolAddress, err)
			continue
		}

		event, err := pegKeeperContract.ParseProfit(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		amount := utils.Quo(event.LpAmount, 18)
		lpPrice := utils.GetHistoricalPriceTokenPrice(poolAddress, "ethereum", uint64(time.Now().Unix()))
		if lpPrice == 0 {
			lpPrice = 1
		}

		feesChan <- interfaces.PegKeeperFee{
			PegKeeperAddress: pegKeeperAddress,
			Amount:           amount,
			LpPrice:          lpPrice,
			TxHash:           vLog.TxHash,
			BlockNumber:      vLog.BlockNumber,
			BlockTimestamp:   block.Time(),
		}
	}
}

func getVolume(tokenExchangeChan chan interfaces.TokenExchange, client *ethclient.Client, pegKeeperAddress common.Address, poolAddress common.Address, opts *bind.CallOpts, config interfaces.Config) {
	from := config.LastBlock
	if from == 0 {
		from = 17258030
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from) + 1),
		ToBlock:   opts.BlockNumber,
		Addresses: []common.Address{poolAddress},
		Topics:    [][]common.Hash{{common.HexToHash("0x8b3e96f2b889fa771c53c981b40daf005f63f637f1869f707052d15a3dd97140")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return
	}

	pegKeeperPoolContract, err := pegKeeperPool.NewPegKeeperPool(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	coins := make(map[*big.Int]common.Address)

	for _, vLog := range logs {

		fmt.Println(pegKeeperAddress, vLog.BlockNumber)

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println("BlockByNumber", poolAddress, err)
			continue
		}

		event, err := pegKeeperPoolContract.ParseTokenExchange(vLog)
		if err != nil {
			fmt.Println("ParseTokenExchange", poolAddress, err)
			continue
		}

		tokenSoldAddress, exists := coins[event.TokensSold]
		if !exists {
			tokenSoldAddress, err = pegKeeperPoolContract.Coins(nil, event.SoldId)
			if err != nil {
				fmt.Println("Coins TokensSold", poolAddress, vLog.TxHash, event.TokensSold, err)
				continue
			}

			coins[event.TokensSold] = tokenSoldAddress
		}

		amount := utils.Quo(event.TokensSold, 18)
		lpPrice := utils.GetHistoricalPriceTokenPrice(tokenSoldAddress, "ethereum", block.Time())
		if lpPrice == 0 {
			lpPrice = 1
		}

		tokenExchangeChan <- interfaces.TokenExchange{
			PegKeeperAddress: pegKeeperAddress,
			Amount:           amount,
			LpPrice:          lpPrice,
			TxHash:           vLog.TxHash,
			BlockNumber:      vLog.BlockNumber,
			BlockTimestamp:   block.Time(),
		}
	}
}

func writePegKeepers(pegKeepers []interfaces.PegKeeper) {
	file, err := json.Marshal(pegKeepers)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PEGKEEPER_DATA, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readPegKeepers() []interfaces.PegKeeper {

	if !utils.FileExists(PEGKEEPER_DATA) {
		return make([]interfaces.PegKeeper, 0)
	}

	file, err := os.ReadFile(PEGKEEPER_DATA)
	if err != nil {
		log.Fatal(err)
	}

	pegkepeers := make([]interfaces.PegKeeper, 0)
	if err := json.Unmarshal([]byte(file), &pegkepeers); err != nil {
		log.Fatal(err)
	}

	return pegkepeers
}

func writePegKeepersFees(fees []interfaces.PegKeeperFee) {
	file, err := json.Marshal(fees)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PEGKEEPER_FEE_DATA, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readPegKeepersFees() []interfaces.PegKeeperFee {

	if !utils.FileExists(PEGKEEPER_FEE_DATA) {
		return make([]interfaces.PegKeeperFee, 0)
	}

	file, err := os.ReadFile(PEGKEEPER_FEE_DATA)
	if err != nil {
		log.Fatal(err)
	}

	fees := make([]interfaces.PegKeeperFee, 0)
	if err := json.Unmarshal([]byte(file), &fees); err != nil {
		log.Fatal(err)
	}

	return fees
}

// DO NOT USE
// Was a function to generate historical values
func IntegrateHistoricalData() {

	pegKeepers := readPegKeepers()

	file, err := os.ReadFile("./volume.json")
	if err != nil {
		log.Fatal(err)
	}

	volumes := make([][]string, 0)
	if err := json.Unmarshal([]byte(file), &volumes); err != nil {
		log.Fatal(err)
	}

	file, err = os.ReadFile("./fee.json")
	if err != nil {
		log.Fatal(err)
	}

	fees := make([][]string, 0)
	if err := json.Unmarshal([]byte(file), &fees); err != nil {
		log.Fatal(err)
	}

	realVolumes := make([]interfaces.TokenExchange, 0)
	for _, volume := range volumes {
		amout, success := new(big.Int).SetString(volume[1], 10)
		if !success {
			log.Fatal("amount volume")
		}

		pegKeeperAddress := common.HexToAddress(volume[0])

		for _, peg := range pegKeepers {
			if strings.EqualFold(peg.PoolAddress.Hex(), volume[0]) {
				pegKeeperAddress = peg.Address
				break
			}
		}

		blockNumber, err := strconv.Atoi(volume[3])
		if err != nil {
			log.Fatal(err)
		}

		blockTimestamp, err := strconv.Atoi(volume[2])
		if err != nil {
			log.Fatal(err)
		}

		amountF := utils.Quo(amout, 18)
		realVolumes = append(realVolumes, interfaces.TokenExchange{
			PegKeeperAddress: pegKeeperAddress,
			Amount:           amountF,
			LpPrice:          1,
			TxHash:           common.HexToHash("0xd436047afadb45afcde419135b8a5655f405f834cbbb25cf5144ab84792dc4c5"),
			BlockNumber:      uint64(blockNumber),
			BlockTimestamp:   uint64(blockTimestamp),
		})
	}

	realFees := make([]interfaces.PegKeeperFee, 0)
	for _, fee := range fees {
		amout, success := new(big.Int).SetString(fee[1], 10)
		if !success {
			log.Fatal("amount volume")
		}

		blockNumber, err := strconv.Atoi(fee[3])
		if err != nil {
			log.Fatal(err)
		}

		blockTimestamp, err := strconv.Atoi(fee[2])
		if err != nil {
			log.Fatal(err)
		}

		amountF := utils.Quo(amout, 18)
		realFees = append(realFees, interfaces.PegKeeperFee{
			PegKeeperAddress: common.HexToAddress(fee[0]),
			Amount:           amountF,
			LpPrice:          1,
			TxHash:           common.HexToHash("0xd436047afadb45afcde419135b8a5655f405f834cbbb25cf5144ab84792dc4c5"),
			BlockNumber:      uint64(blockNumber),
			BlockTimestamp:   uint64(blockTimestamp),
		})
	}

	fileFee, err := json.Marshal(realFees)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PEGKEEPER_FEE_DATA, fileFee, 0644); err != nil {
		log.Fatal(err)
	}

	fileVolume, err := json.Marshal(realVolumes)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PEGKEEPER_VOLUME_DATA, fileVolume, 0644); err != nil {
		log.Fatal(err)
	}
}
