package src

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/contracts/curvePoolFactoryV2"
	"main/contracts/erc20"
	"main/contracts/feeCollector"
	"main/contracts/poolOwnerPendingFees"
	"main/interfaces"
	"main/utils"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	functionSignature                  = "withdraw_admin_fees()"
	PENDING_POOL_FEES_DATA             = "./data/pendingPoolFees/data.json"
	POOL_WITH_WITHDRAW_ADMIN_FEES_DATA = "./data/pendingPoolFees/pool-with-withdraw-admin-fees.json"
	BURNERS_FOR_TOKEN_DATA             = "./data/pendingPoolFees/burnes-for-tokens.json"
	BURNERS_DATA                       = "./data/pendingPoolFees/burners.json"
	COWSWAP_FEES_PER_TOKEN_PATH        = "./data/pendingPoolFees/cowswap-fees-per-token"
	NB_ADDRESSES_IN_FUNCTION           = 20
	ETHEREUM_CHAIN                     = "ethereum"
	POLYGON_CHAIN                      = "polygon"
	ARBITRUM_CHAIN                     = "arbitrum"
	OPTIMISM_CHAIN                     = "optimism"
	AVALANCHE_CHAIN                    = "avalanche"
	FANTOM_CHAIN                       = "fantom"
)

var (
	CHAINS_APPROVED       = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN, AVALANCHE_CHAIN, FANTOM_CHAIN}
	COWSWAP_BURNER_CHAINS = []string{ETHEREUM_CHAIN}
	NODE_TENDERLY         = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN}
	FORK_CHAIN_ID         = map[string]string{
		AVALANCHE_CHAIN: "43114",
		FANTOM_CHAIN:    "250",
	}
)

var COWSWAP_CHAINS = map[string]string{
	ETHEREUM_CHAIN: "mainnet",
	ARBITRUM_CHAIN: "arbitrum_one",
}

// Contract which receive fee tokens + burn + withdraw
var POOL_OWNERS = map[string]common.Address{
	ETHEREUM_CHAIN:  utils.POOL_OWNER_PENDING_FEES_MAINNET,
	POLYGON_CHAIN:   utils.ADMIN_RECEIVER_FEE_POLYGON,
	OPTIMISM_CHAIN:  utils.ADMIN_RECEIVER_FEE_OPTIMISM,
	ARBITRUM_CHAIN:  utils.ADMIN_RECEIVER_FEE_ARBITRUM,
	AVALANCHE_CHAIN: utils.ADMIN_RECEIVER_FEE_AVALANCHE,
	FANTOM_CHAIN:    utils.ADMIN_RECEIVER_FEE_FANTOM,
}

// Contract which receive target token to distribute
var FEE_RECEIVERS = map[string]common.Address{
	POLYGON_CHAIN:   utils.ADMIN_RECEIVER_FEE_POLYGON,
	OPTIMISM_CHAIN:  utils.ADMIN_RECEIVER_FEE_OPTIMISM,
	ARBITRUM_CHAIN:  utils.ADMIN_RECEIVER_FEE_ARBITRUM,
	AVALANCHE_CHAIN: utils.ADMIN_RECEIVER_FEE_AVALANCHE,
	FANTOM_CHAIN:    utils.ADMIN_RECEIVER_FEE_FANTOM,
}

// Tokens + LP which will be burn to the final target token
var TARGET_TOKEN_BEFORE_SEND = map[string][]common.Address{
	ETHEREUM_CHAIN: {
		utils.USDC,
		utils.DAI,
		utils.USDT,
		utils.THREE_CRV,
	},
	POLYGON_CHAIN:   {utils.LP_TARGET_TOKEN_POLYGON},
	OPTIMISM_CHAIN:  {utils.LP_TARGET_TOKEN_OPTIMISM},
	ARBITRUM_CHAIN:  {utils.LP_TARGET_TOKEN_ARBITRUM},
	AVALANCHE_CHAIN: {utils.TARGET_TOKEN_AVALANCHE},
	FANTOM_CHAIN:    {utils.TARGET_TOKEN_FANTOM},
}

// Target token is the token which will be distributed / bridged
var TARGET_TOKEN = map[string]common.Address{
	ETHEREUM_CHAIN:  utils.THREE_CRV,
	POLYGON_CHAIN:   utils.TARGET_TOKEN_POLYGON,
	OPTIMISM_CHAIN:  utils.TARGET_TOKEN_OPTIMISM,
	ARBITRUM_CHAIN:  utils.TARGET_TOKEN_ARBITRUM,
	AVALANCHE_CHAIN: utils.TARGET_TOKEN_AVALANCHE,
	FANTOM_CHAIN:    utils.TARGET_TOKEN_FANTOM,
}

var TARKET_TOKEN_DECIMALS = map[string]uint64{
	ETHEREUM_CHAIN:  18,
	POLYGON_CHAIN:   6,
	OPTIMISM_CHAIN:  18,
	ARBITRUM_CHAIN:  18,
	AVALANCHE_CHAIN: 18,
	FANTOM_CHAIN:    6,
}

// Fee collector per chain with cowswap burner
var FEE_COLLECTOR_PER_CHAIN_COWSWAP_BURNER = map[string]common.Address{
	ETHEREUM_CHAIN: utils.FEE_COLLECTOR_MAINNET,
}

func PendingPoolFees(client *ethclient.Client, currentBlockf uint64, currentBlockTImestamp uint64) {
	currentPendingCrvFees := readPending3CRVFees()

	// We retrieve all the pools from the Curve API
	allPools, err := utils.GetAllCurvePools()
	if err != nil {
		log.Fatal(err)
	}

	earned := make(map[string]float64)

	for _, chain := range CHAINS_APPROVED {
		fmt.Println("computing fees for chain ", chain)
		if utils.ArrayContains(COWSWAP_BURNER_CHAINS, chain) {
			earned[chain] = estimatedWithCowswapBurner(client, allPools, chain, COWSWAP_CHAINS[chain])
		} else {
			earned[chain] = estimatedWithSimulation(client, allPools, chain)
		}
	}

	fmt.Println(earned)

	lastDistributionTimestamp, lastDistributionAmount := fetchLastDistribution(client, currentBlockf, currentPendingCrvFees.LastBlock, currentPendingCrvFees.LastDistributionTimestamp, currentPendingCrvFees.LastDistributionAmount)
	writePending3CRVFees(interfaces.PendingPoolFees{
		LastBlock:                 currentBlockf,
		LastDistributionTimestamp: lastDistributionTimestamp,
		LastDistributionAmount:    lastDistributionAmount,
		Pending3CRVFees:           earned,
	})
}

func estimatedWithCowswapBurner(mainnetClient *ethclient.Client, allPools []interfaces.CurvePool, chain string, cowswapChain string) float64 {

	// Get fee collector address
	feeCollectorAddress, exists := FEE_COLLECTOR_PER_CHAIN_COWSWAP_BURNER[chain]
	if !exists {
		return 0
	}

	client := mainnetClient
	var err error
	if !strings.EqualFold(chain, ETHEREUM_CHAIN) {
		rpcUrl := utils.GetPublicRpcUrl(chain)
		if len(rpcUrl) == 0 {
			return 0
		}

		client, err = ethclient.Dial(rpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", rpcUrl)
			return 0
		}
	}

	// Get burner
	feeCollectorContract, err := feeCollector.NewFeeCollector(feeCollectorAddress, client)
	if err != nil {
		return 0
	}

	burner, err := feeCollectorContract.Burner(nil)
	if err != nil {
		return 0
	}

	balanceOfAddresses := []common.Address{burner, feeCollectorAddress}

	// For all pool tokens + lp, check the fee collector balance
	// Generate calldatas
	coinsToQuotePerContract := make(map[common.Address]map[common.Address]*big.Int)

	for _, balanceOfAddress := range balanceOfAddresses {
		_, exists := coinsToQuotePerContract[balanceOfAddress]
		if !exists {
			coinsToQuotePerContract[balanceOfAddress] = make(map[common.Address]*big.Int, 0)
		}

		for _, pool := range allPools {
			if !strings.EqualFold(pool.BlockchainId, chain) {
				continue
			}

			coin := common.HexToAddress(pool.LpTokenAddress)

			_, exists := coinsToQuotePerContract[balanceOfAddress][coin]
			if !exists {
				erc20Contract, err := erc20.NewErc20(coin, client)
				if err != nil {
					fmt.Println(pool.LpTokenAddress, err)
					continue
				}

				balanceOf, err := erc20Contract.BalanceOf(nil, balanceOfAddress)
				if err != nil {
					fmt.Println(pool.LpTokenAddress, err)
					continue
				}
				coinsToQuotePerContract[balanceOfAddress][coin] = balanceOf
			}

			for _, coin := range pool.Coins {

				_, exists := coinsToQuotePerContract[balanceOfAddress][common.HexToAddress(coin.Address)]
				if !exists {

					erc20Contract, err := erc20.NewErc20(common.HexToAddress(coin.Address), client)
					if err != nil {
						fmt.Println(coin.Address, err)
						continue
					}

					balanceOf, err := erc20Contract.BalanceOf(nil, balanceOfAddress)
					if err != nil {
						fmt.Println(coin.Address, err)
						continue
					}

					coinsToQuotePerContract[balanceOfAddress][common.HexToAddress(coin.Address)] = balanceOf
				}
			}
		}
	}

	// Check withdraw_admin_fees
	poolsWithWithdrawAdminFee := getPoolsWithWithdrawAdminFee(client, allPools, chain)
	for _, pool := range poolsWithWithdrawAdminFee {
		poolContract, err := curvePoolFactoryV2.NewCurvePoolFactoryV2(pool, client)
		if err != nil {
			fmt.Println(err)
			continue
		}

		i := big.NewInt(0)
		for {
			coin, err := poolContract.Coins(nil, i)
			if err != nil {
				break
			}

			balance, err := poolContract.AdminBalances(nil, i)
			if err != nil {
				break
			}

			value, exists := coinsToQuotePerContract[feeCollectorAddress][coin]
			if !exists {
				coinsToQuotePerContract[feeCollectorAddress][coin] = balance
			} else {
				coinsToQuotePerContract[feeCollectorAddress][coin] = new(big.Int).Add(value, balance)
			}

			i = new(big.Int).Add(i, big.NewInt(1))
		}
	}

	// Get Cowswap quotes from cow api
	fmt.Println("Quotes to do : ", len(coinsToQuotePerContract))
	totalFees := 0.0
	feesPerToken := make(map[string]float64, 0)

	balances := make(map[common.Address]*big.Int)
	for _, quote := range coinsToQuotePerContract {
		for token, balance := range quote {
			value, exists := balances[token]
			if !exists {
				balances[token] = balance
			} else {
				balances[token] = new(big.Int).Add(value, balance)
			}
		}
	}

	for coin, balance := range balances {
		if balance.Cmp(big.NewInt(0)) == 0 {
			// Empty balance
			continue
		}

		quoteResp, err := quote(cowswapChain, coin, balance)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(quoteResp.Quote.BuyAmount) == 0 {
			for _, pool := range allPools {
				found := false
				for _, coinPool := range pool.Coins {
					if strings.EqualFold(coin.Hex(), coinPool.Address) {
						found = true

						erc20Contract, err := erc20.NewErc20(coin, client)
						if err != nil {
							break
						}

						decimals, err := erc20Contract.Decimals(nil)
						if err != nil {
							break
						}

						if coinPool.UsdPrice == nil {
							coinPool.UsdPrice = 0
						}
						amount := coinPool.UsdPrice.(float64) * utils.Quo(balance, uint64(decimals))
						totalFees += amount
						feesPerToken[coin.Hex()] = amount
						break
					}
				}

				if found {
					break
				}
			}
			continue
		} else {
			bigIntValue := new(big.Int)

			// Convertir la chaîne de caractères en big.Int
			bigIntValue, success := bigIntValue.SetString(quoteResp.Quote.BuyAmount, 10)
			if !success {
				fmt.Println("Erreur de conversion de la chaîne en big.Int")
				continue
			}

			amount := utils.Quo(bigIntValue, 18)
			totalFees += amount
			feesPerToken[coin.Hex()] = amount
		}
	}

	writeMapFees(feesPerToken, COWSWAP_FEES_PER_TOKEN_PATH+"-"+chain+".json")
	fmt.Println("totalFees", totalFees)
	return totalFees
}

func quote(chain string, tokenToSell common.Address, amountToSell *big.Int) (*interfaces.CowswapQuote, error) {
	posturl := "https://api.cow.fi/" + chain + "/api/v1/quote"

	body := []byte(`{
		"sellToken": "` + tokenToSell.Hex() + `",
		"buyToken": "0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E",
		"receiver": "0xD16d5eC345Dd86Fb63C6a9C43c517210F1027914",
		"sellTokenBalance": "erc20",
		"buyTokenBalance": "erc20",
		"from": "0xD16d5eC345Dd86Fb63C6a9C43c517210F1027914",
		"priceQuality": "verified",
		"signingScheme": "eip712",
		"onchainOrder": false,
		"kind": "sell",
		"sellAmountBeforeFee": "` + amountToSell.String() + `"
	}`)

	fmt.Println(`{
		"sellToken": "` + tokenToSell.Hex() + `",
		"buyToken": "0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E",
		"receiver": "0xD16d5eC345Dd86Fb63C6a9C43c517210F1027914",
		"sellTokenBalance": "erc20",
		"buyTokenBalance": "erc20",
		"from": "0xD16d5eC345Dd86Fb63C6a9C43c517210F1027914",
		"priceQuality": "verified",
		"signingScheme": "eip712",
		"onchainOrder": false,
		"kind": "sell",
		"sellAmountBeforeFee": "` + amountToSell.String() + `"
	}`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	quote := new(interfaces.CowswapQuote)

	err = json.Unmarshal([]byte(body), &quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func estimatedWithSimulation(client *ethclient.Client, allPools []interfaces.CurvePool, chain string) float64 {
	pools := make([]interfaces.CurvePool, 0)
	for _, pool := range allPools {
		if strings.EqualFold(chain, pool.BlockchainId) {
			pools = append(pools, pool)
		}
	}

	clientChain := client
	var err error
	if !strings.EqualFold(chain, ETHEREUM_CHAIN) {
		rpcUrl := utils.GetPublicRpcUrl(chain)
		if len(rpcUrl) == 0 {
			return 0
		}

		clientChain, err = ethclient.Dial(rpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", rpcUrl)
			return 0
		}
	}

	// Generate calldatas for withdraw_many
	poolsWithWithdrawAdminFee := getPoolsWithWithdrawAdminFee(clientChain, pools, chain)
	// 0x3165fbc056036ad3bc38ddf39f1764220a0e562a
	//poolsWithWithdrawAdminFee = removeRevert(clientChain, "withdraw_admin_fees", poolsWithWithdrawAdminFee, chain)
	//chunkPoolsWithdrawAdminFees := utils.ChunkAddressSlice(poolsWithWithdrawAdminFee, NB_ADDRESSES_IN_FUNCTION)
	calldatasWithdrawFees := getCalldatas("withdraw_admin_fees", poolsWithWithdrawAdminFee)

	// Generate calldatas for burn_many
	coinsWithBurnersSimpe := getTokensWithBurners(clientChain, pools, chain)
	//coinsWithBurnersSimpe = removeRevert(clientChain, "burn", coinsWithBurnersSimpe, chain)

	coinsWithBurnersSimpe = manageTokens(coinsWithBurnersSimpe, TARGET_TOKEN_BEFORE_SEND[chain], pools, chain)

	//coinsWithBurners := utils.ChunkAddressSlice(coinsWithBurnersSimpe, NB_ADDRESSES_IN_FUNCTION)
	calldatasBurnMany := getCalldatas("burn", coinsWithBurnersSimpe)

	// Generate target token balance calldata
	targetTokenBalanceRequest := getTargetTokenBalanceCallData(chain)

	balance := big.NewInt(0)
	newBalance := big.NewInt(0)

	if utils.ArrayContains(NODE_TENDERLY, chain) {
		// Use Tenderly node

		// Send simulation
		simulationResponse := sendSimulation(chain, calldatasWithdrawFees, calldatasBurnMany, targetTokenBalanceRequest)

		// Extract 3CRV balance
		firstResult := simulationResponse.Result[0]
		lastResult := simulationResponse.Result[len(simulationResponse.Result)-1]

		// Extract data
		var success bool
		balance, success = new(big.Int).SetString(firstResult.Trace[0].DecodedOutput[0].Value.(string), 10)
		if !success {
			log.Fatal(err)
		}

		newBalance, success = new(big.Int).SetString(lastResult.Trace[0].DecodedOutput[0].Value.(string), 10)
		if !success {
			log.Fatal(err)
		}
	} else {
		// Use fork
		forkRpcUrl, forkId, _, err := createFork(FORK_CHAIN_ID[chain], 0)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		defer deleteFork(forkId)

		forkClient, err := ethclient.Dial(forkRpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", forkRpcUrl)
			return 0
		}

		// Simulation de l'exécution de la transaction
		targetTokenContract, err := erc20.NewErc20(TARGET_TOKEN[chain], forkClient)
		if err != nil {
			log.Fatal(err)
		}

		balance, err = targetTokenContract.BalanceOf(nil, FEE_RECEIVERS[chain])

		to := POOL_OWNERS[chain]

		for _, call := range calldatasWithdrawFees {
			sendTx(forkClient, to, call)
		}

		for _, call := range calldatasBurnMany {
			sendTx(forkClient, to, call)
		}

		newBalance, err = targetTokenContract.BalanceOf(nil, FEE_RECEIVERS[chain])
	}

	return utils.Quo(new(big.Int).Sub(newBalance, balance), TARKET_TOKEN_DECIMALS[chain])
}

func fetchLastDistribution(client *ethclient.Client, currentBlock uint64, lastBlock uint64, lastTimestamp uint64, lastDistributionAmount uint64) (uint64, uint64) {

	from := lastBlock
	if from == 0 {
		from = 19526061
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from) + 1),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{utils.CRVUSD_ADDRESS},
		Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return lastTimestamp, lastDistributionAmount
	}

	timestamp := lastTimestamp
	for _, vLog := range logs {
		contract, err := erc20.NewErc20(vLog.Address, client)
		if err != nil {
			fmt.Println(err)
			continue
		}

		event, err := contract.ParseTransfer(vLog)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !strings.EqualFold(event.To.Hex(), utils.FEE_DISTRIBUTOR_MAINNET.Hex()) {
			continue
		}

		if !strings.EqualFold(utils.HOOKER_MAINNET.Hex(), event.From.Hex()) {
			continue
		}

		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			fmt.Println(err)
			continue
		}

		if block.Time()-timestamp > 2*utils.DAY_TO_SEC {
			lastDistributionAmount = 0
		}
		timestamp = block.Time()
		lastDistributionAmount += uint64(utils.Quo(event.Value, 18))
	}

	return timestamp, lastDistributionAmount

}

func getPoolsWithWithdrawAdminFee(client *ethclient.Client, pools []interfaces.CurvePool, chain string) []common.Address {

	poolWithWithdrawAdminFee := readMapBool(POOL_WITH_WITHDRAW_ADMIN_FEES_DATA)

	// Function selector
	selector := crypto.Keccak256Hash([]byte(functionSignature)).Hex()[2:10]

	for _, pool := range pools {

		_, exists := poolWithWithdrawAdminFee[chain]
		if !exists {
			poolWithWithdrawAdminFee[chain] = make(map[common.Address]bool)
		}

		poolAddress := common.HexToAddress(pool.Address)
		_, exists = poolWithWithdrawAdminFee[chain][poolAddress]
		if exists {
			continue
		}

		bytecode, err := client.CodeAt(context.Background(), poolAddress, nil) // nil pour le dernier bloc connu
		if err != nil {
			log.Fatalf("Erreur lors de la récupération du bytecode : %s", err)
		}

		poolWithWithdrawAdminFee[chain][poolAddress] = strings.Contains(common.Bytes2Hex(bytecode), selector)
	}

	writeMapBol(poolWithWithdrawAdminFee, POOL_WITH_WITHDRAW_ADMIN_FEES_DATA)

	response := make([]common.Address, 0)
	for poolAddress, value := range poolWithWithdrawAdminFee[chain] {
		if value {
			response = append(response, poolAddress)
		}
	}
	return response
}

func getTenderlyCallDatas(chain string, calldatas [][]byte) []interfaces.TenderlyCallData {

	tenderlyCallDatas := make([]interfaces.TenderlyCallData, 0)

	for _, calldata := range calldatas {
		tenderlyCallDatas = append(tenderlyCallDatas, interfaces.TenderlyCallData{
			From:           utils.POOL_OWNER_PENDING_FEES_OWNER,
			To:             POOL_OWNERS[chain],
			Data:           "0x" + common.Bytes2Hex(calldata),
			SimulationType: "quick",
		})
	}

	return tenderlyCallDatas
}

func getTokensWithBurners(client *ethclient.Client, pools []interfaces.CurvePool, chain string) []common.Address {
	poolOwnerContract, err := poolOwnerPendingFees.NewPoolOwnerPendingFees(POOL_OWNERS[chain], client)
	if err != nil {
		log.Fatal(err)
	}

	tokensWithBurners := readMapBool(BURNERS_FOR_TOKEN_DATA)
	burners := readMap(BURNERS_DATA)

	for _, pool := range pools {

		_, exists := tokensWithBurners[chain]
		if !exists {
			tokensWithBurners[chain] = make(map[common.Address]bool)
		}

		lpAddress := common.HexToAddress(pool.LpTokenAddress)
		haveBurner, exists := tokensWithBurners[chain][lpAddress]
		if !exists || !haveBurner {
			burnerAddress, err := poolOwnerContract.Burners(nil, lpAddress)
			if err == nil {
				tokensWithBurners[chain][lpAddress] = !utils.IsNullAddress(burnerAddress)
			}
		}

		allCoins := make([]interfaces.Coin, 0)
		allCoins = append(allCoins, pool.Coins...)
		allCoins = append(allCoins, pool.UnderlyingCoins...)

		for _, coin := range allCoins {
			coinAddress := common.HexToAddress(coin.Address)
			haveBurner, exists := tokensWithBurners[chain][coinAddress]
			if exists && haveBurner {
				continue
			}

			burnerAddress, err := poolOwnerContract.Burners(nil, coinAddress)
			if err != nil {
				continue
			}

			tokensWithBurners[chain][coinAddress] = !utils.IsNullAddress(burnerAddress)
			burners[burnerAddress] = true
		}
	}

	writeMapBol(tokensWithBurners, BURNERS_FOR_TOKEN_DATA)
	writeMap(burners, BURNERS_DATA)

	response := make([]common.Address, 0)
	for coinAddress, value := range tokensWithBurners[chain] {
		if value {
			response = append(response, coinAddress)
		}
	}
	return response
}

func getCalldatas(functionName string, addressesChunk []common.Address) [][]byte {
	abi, err := poolOwnerPendingFees.PoolOwnerPendingFeesMetaData.GetAbi()
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	calldatas := make([][]byte, 0)
	for _, chunk := range addressesChunk {

		calldata, err := abi.Pack(functionName, chunk)
		if err != nil {
			log.Fatalf("Failed to generate calldata: %v", err)
		}

		calldatas = append(calldatas, calldata)
	}

	return calldatas
}

func removeRevert(client *ethclient.Client, functionName string, addrs []common.Address, chain string) []common.Address {
	abi, err := poolOwnerPendingFees.PoolOwnerPendingFeesMetaData.GetAbi()
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	response := make([]common.Address, 0)
	for _, addr := range addrs {

		calldata, err := abi.Pack(functionName, addr)
		if err != nil {
			log.Fatalf("Failed to generate calldata: %v", err)
		}

		if callRevert(client, chain, calldata) {
			continue
		}

		response = append(response, addr)
	}

	return response
}

func getTargetTokenBalanceCallData(chain string) interfaces.TenderlyCallData {
	abi, err := erc20.Erc20MetaData.GetAbi()
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	calldata, err := abi.Pack("balanceOf", FEE_RECEIVERS[chain])
	if err != nil {
		log.Fatalf("Failed to generate calldata: %v", err)
	}

	return interfaces.TenderlyCallData{
		From:           utils.NULL_ADDRESS,
		To:             TARGET_TOKEN[chain],
		Data:           "0x" + common.Bytes2Hex(calldata),
		SimulationType: "quick",
	}
}

func sendSimulation(chain string, calldatasWithdrawFees [][]byte, calldatasBurnMany [][]byte, threeCRVBalanceRequest interfaces.TenderlyCallData) *interfaces.TenderlyBatchSimulationResponse {
	// Send Tenderly simulation
	tenderlyChain := chain
	if strings.EqualFold(chain, ETHEREUM_CHAIN) {
		tenderlyChain = "mainnet"
	}
	posturl := "https://" + tenderlyChain + ".gateway.tenderly.co/" + utils.GoDotEnvVariable("TENDERLY_NODE")

	calldataParams := make([]interface{}, 0)
	calldataParams = append(calldataParams, threeCRVBalanceRequest)

	for _, t := range getTenderlyCallDatas(chain, calldatasWithdrawFees) {
		calldataParams = append(calldataParams, t)
	}

	for _, t := range getTenderlyCallDatas(chain, calldatasBurnMany) {
		calldataParams = append(calldataParams, t)
	}

	calldataParams = append(calldataParams, threeCRVBalanceRequest)

	//fmt.Println(len(calldataParams))
	params := make([]interface{}, 0)
	params = append(params, calldataParams)
	params = append(params, "latest")

	body, err := json.Marshal(interfaces.TenderlyBatchSimulation{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "tenderly_simulateBundle",
		Params:  params,
	})
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	simulationResponse := &interfaces.TenderlyBatchSimulationResponse{}
	derr := json.NewDecoder(res.Body).Decode(simulationResponse)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	return simulationResponse
}

func callRevert(client *ethclient.Client, chain string, calldata []byte) bool {
	// Construction de la transaction d'appel
	to := POOL_OWNERS[chain]
	msg := ethereum.CallMsg{
		To:   &to,
		Data: calldata,
	}

	// Simulation de l'exécution de la transaction
	_, err := client.CallContract(context.Background(), msg, nil)
	return err != nil
}

func manageTokens(coins []common.Address, coinsToManage []common.Address, allPools []interfaces.CurvePool, chain string) []common.Address {
	coinsWithBurnersSimpeSlice := make([]common.Address, 0)
	coinSkipped := make([]common.Address, 0)

	// Remove coins to put at the end
	for _, coin := range coins {
		found := false
		for _, coinToManage := range coinsToManage {
			if strings.EqualFold(coin.Hex(), coinToManage.Hex()) {
				coinSkipped = append(coinSkipped, coinToManage)
				found = true
				break
			}
		}

		if !found {
			coinsWithBurnersSimpeSlice = append(coinsWithBurnersSimpeSlice, coin)
		}
	}

	// Order tokens, LP first, token after
	lpTokens := make(map[common.Address]bool)
	tokens := make(map[common.Address]bool)

	for _, coin := range coinsWithBurnersSimpeSlice {
		for _, pool := range allPools {

			if strings.EqualFold(pool.LpTokenAddress, coin.Hex()) {
				lpTokens[coin] = true
				break
			}

			allCoins := make([]interfaces.Coin, 0)
			allCoins = append(allCoins, pool.Coins...)
			allCoins = append(allCoins, pool.UnderlyingCoins...)

			found := false
			for _, c := range allCoins {
				if strings.EqualFold(c.Address, coin.Hex()) {
					tokens[coin] = true
					found = true
					break
				}
			}

			if found {
				break
			}
		}
	}

	results := make([]common.Address, 0)
	for coin := range lpTokens {
		results = append(results, coin)
	}
	for coin := range tokens {
		results = append(results, coin)
	}
	for coin := range tokens {
		results = append(results, coin)
	}

	for _, coinToManage := range coinsToManage {
		for _, cs := range coinSkipped {
			if strings.EqualFold(cs.Hex(), coinToManage.Hex()) {
				results = append(results, coinToManage)
				break
			}
		}
	}

	return results
}

func createFork(chainId string, blockNumber int64) (string, string, *ecdsa.PrivateKey, error) {
	var forkRequest interfaces.TenderlyForkRequest

	if blockNumber == 0 {
		forkRequest = interfaces.TenderlyForkRequest{
			NetworkId: chainId,
			Shared:    true,
		}
	} else {
		forkRequest = interfaces.TenderlyForkRequest{
			NetworkId:   chainId,
			BlockNumber: blockNumber,
			Shared:      true,
		}
	}

	body, err := json.Marshal(forkRequest)

	if err != nil {
		return "", "", nil, err
	}

	r, err := http.NewRequest("POST", "https://api.tenderly.co/api/v1/account/"+utils.GoDotEnvVariable("TENDERLY_SLUG")+"/project/"+utils.GoDotEnvVariable("TENDERLY_PROJECT_SLUG")+"/fork", bytes.NewBuffer(body))
	if err != nil {
		return "", "", nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Key", utils.GoDotEnvVariable("TENDERLY_ACCESS_KEY"))

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		return "", "", nil, err
	}

	defer res.Body.Close()

	forkResponse := &interfaces.TenderlyForkResponse{}
	derr := json.NewDecoder(res.Body).Decode(forkResponse)
	if derr != nil {
		return "", "", nil, err
	}

	forkRpcUrl := forkResponse.SimulationFork.RpcUrl

	// Fund account
	privateKey, err := crypto.HexToECDSA(utils.GoDotEnvVariable("FORK_PK"))
	if err != nil {
		return "", "", nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	addEth(httpClient, fromAddress, forkRpcUrl)
	addCrv(httpClient, fromAddress, forkRpcUrl)

	return forkRpcUrl, forkResponse.SimulationFork.Id, privateKey, nil
}

func addEth(httpClient *http.Client, fromAddress common.Address, forkRpcUrl string) error {
	params := make([]interface{}, 0)
	addressesToFund := make([]string, 0)
	addressesToFund = append(addressesToFund, fromAddress.Hex())

	params = append(params, addressesToFund)
	params = append(params, "0x3635c9adc5dea00000")
	body, err := json.Marshal(interfaces.TenderlyAddBalanceRequest{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "tenderly_addBalance",
		Params:  params,
	})
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", forkRpcUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}

func addCrv(httpClient *http.Client, fromAddress common.Address, forkRpcUrl string) error {
	params := make([]interface{}, 0)

	params = append(params, utils.CRV.Hex())
	params = append(params, fromAddress.Hex())
	params = append(params, "0x204FCE5E3E25026110000000")
	body, err := json.Marshal(interfaces.TenderlyAddBalanceRequest{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "tenderly_setErc20Balance",
		Params:  params,
	})
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", forkRpcUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}

func mineBlocks(httpClient *http.Client, forkRpcUrl string, nbBlocks int64) error {
	params := make([]interface{}, 0)

	params = append(params, nbBlocks)
	body, err := json.Marshal(interfaces.TenderlyAddBalanceRequest{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "evm_increaseBlocks",
		Params:  params,
	})
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", forkRpcUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}

func addTime(httpClient *http.Client, forkRpcUrl string, time int64) error {
	params := make([]interface{}, 0)

	params = append(params, time)
	body, err := json.Marshal(interfaces.TenderlyAddBalanceRequest{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "evm_increaseTime",
		Params:  params,
	})
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", forkRpcUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}

func deleteFork(forkId string) {
	r, err := http.NewRequest("DELETE", "https://api.tenderly.co/api/v1/account/"+utils.GoDotEnvVariable("TENDERLY_SLUG")+"/project/"+utils.GoDotEnvVariable("TENDERLY_PROJECT_SLUG")+"/fork/"+forkId, nil)
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}

func sendTx(forkClient *ethclient.Client, to common.Address, data []byte) error {
	privateKey, err := crypto.HexToECDSA(utils.GoDotEnvVariable("FORK_PK"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := forkClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0)
	gasLimit := uint64(5_010_000) // in units
	gasPrice, err := forkClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := to
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := forkClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return forkClient.SendTransaction(context.Background(), signedTx)
}

func readMapBool(path string) map[string]map[common.Address]bool {
	if !utils.FileExists(path) {
		return make(map[string]map[common.Address]bool)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	pools := make(map[string]map[common.Address]bool)
	if err := json.Unmarshal([]byte(file), &pools); err != nil {
		log.Fatal(err)
	}

	return pools
}

func writeMapBol(pools map[string]map[common.Address]bool, path string) {
	file, err := json.Marshal(pools)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readMap(path string) map[common.Address]bool {
	if !utils.FileExists(path) {
		return make(map[common.Address]bool)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	pools := make(map[common.Address]bool)
	if err := json.Unmarshal([]byte(file), &pools); err != nil {
		log.Fatal(err)
	}

	return pools
}

func writeMap(pools map[common.Address]bool, path string) {
	file, err := json.Marshal(pools)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func writeMapFees(pools map[string]float64, path string) {
	file, err := json.Marshal(pools)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readPending3CRVFees() interfaces.PendingPoolFees {
	if !utils.FileExists(PENDING_POOL_FEES_DATA) {
		return interfaces.PendingPoolFees{
			LastBlock:                 0,
			LastDistributionTimestamp: 0,
			LastDistributionAmount:    0,
		}
	}

	file, err := os.ReadFile(PENDING_POOL_FEES_DATA)
	if err != nil {
		log.Fatal(err)
	}

	pendingPoolFees := new(interfaces.PendingPoolFees)
	if err := json.Unmarshal([]byte(file), &pendingPoolFees); err != nil {
		log.Fatal(err)
	}

	return *pendingPoolFees
}

func writePending3CRVFees(data interfaces.PendingPoolFees) {
	file, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(PENDING_POOL_FEES_DATA, file, 0644); err != nil {
		log.Fatal(err)
	}
}
