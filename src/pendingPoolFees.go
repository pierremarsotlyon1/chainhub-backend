package src

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/erc20"
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
	NB_ADDRESSES_IN_FUNCTION           = 20
	ETHEREUM_CHAIN                     = "ethereum"
	POLYGON_CHAIN                      = "polygon"
	ARBITRUM_CHAIN                     = "arbitrum"
	OPTIMISM_CHAIN                     = "optimism"
	AVALANCHE_CHAIN                    = "avalanche"
	FANTOM_CHAIN                       = "fantom"
)

var (
	CHAINS_APPROVED = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN, AVALANCHE_CHAIN, FANTOM_CHAIN}
	NODE_TENDERLY   = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN}
	FORK_CHAIN_ID   = map[string]string{
		AVALANCHE_CHAIN: "43114",
		FANTOM_CHAIN:    "250",
	}
)

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
	ETHEREUM_CHAIN:  utils.FEES_DISTRIBUTOR_MAINNET,
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

func PendingPoolFees(client *ethclient.Client, currentBlockf uint64, currentBlockTImestamp uint64) {
	currentPendingCrvFees := readPending3CRVFees()

	// We retrieve all the pools from the Curve API
	allPools, err := utils.GetAllCurvePools()
	if err != nil {
		log.Fatal(err)
	}

	earned := make(map[string]float64)

	for _, chain := range CHAINS_APPROVED {

		pools := make([]interfaces.CurvePool, 0)
		for _, pool := range allPools {
			if strings.EqualFold(chain, pool.BlockchainId) {
				pools = append(pools, pool)
			}
		}

		clientChain := client
		if !strings.EqualFold(chain, ETHEREUM_CHAIN) {
			rpcUrl := utils.GetPublicRpcUrl(chain)
			if len(rpcUrl) == 0 {
				continue
			}

			clientChain, err = ethclient.Dial(rpcUrl)
			if err != nil {
				fmt.Println("Error when creating rpc ", rpcUrl)
				continue
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
			forkRpcUrl, forkId := createFork(FORK_CHAIN_ID[chain])
			defer deleteFork(forkId)

			forkClient, err := ethclient.Dial(forkRpcUrl)
			if err != nil {
				fmt.Println("Error when creating rpc ", forkRpcUrl)
				continue
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

		earned[chain] = utils.Quo(new(big.Int).Sub(newBalance, balance), TARKET_TOKEN_DECIMALS[chain])
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

func fetchLastDistribution(client *ethclient.Client, currentBlock uint64, lastBlock uint64, lastTimestamp uint64, lastDistributionAmount uint64) (uint64, uint64) {

	from := lastBlock
	if from == 0 {
		from = 19526062
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{utils.THREE_CRV},
		Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return lastTimestamp, lastDistributionAmount
	}

	burners := readMap(BURNERS_DATA)

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

		if !strings.EqualFold(event.To.Hex(), utils.FEES_DISTRIBUTOR_MAINNET.Hex()) {
			continue
		}

		from := event.From.Hex()
		ok := strings.EqualFold(utils.POOL_OWNER_PENDING_FEES_MAINNET.Hex(), from) || strings.EqualFold(utils.CURVE_SWAP_ROUTER.Hex(), from)

		if !ok {
			for burnerAddress := range burners {
				if strings.EqualFold(from, burnerAddress.Hex()) {
					ok = true
					break
				}
			}
		}

		if !ok {
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
		_, exists = tokensWithBurners[chain][lpAddress]
		if !exists {
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
			_, exists := tokensWithBurners[chain][coinAddress]
			if exists {
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

func createFork(chainId string) (string, string) {
	body, err := json.Marshal(interfaces.TenderlyForkRequest{
		NetworkId: chainId,
	})

	if err != nil {
		log.Fatal(err)
	}

	r, err := http.NewRequest("POST", "https://api.tenderly.co/api/v1/account/"+utils.GoDotEnvVariable("TENDERLY_SLUG")+"/project/"+utils.GoDotEnvVariable("TENDERLY_PROJECT_SLUG")+"/fork", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Key", utils.GoDotEnvVariable("TENDERLY_ACCESS_KEY"))

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	forkResponse := &interfaces.TenderlyForkResponse{}
	derr := json.NewDecoder(res.Body).Decode(forkResponse)
	if derr != nil {
		panic(derr)
	}

	forkRpcUrl := forkResponse.SimulationFork.RpcUrl

	// Fund account
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

	params := make([]interface{}, 0)
	addressesToFund := make([]string, 0)
	addressesToFund = append(addressesToFund, fromAddress.Hex())

	params = append(params, addressesToFund)
	params = append(params, "0x3635c9adc5dea00000")
	body, err = json.Marshal(interfaces.TenderlyAddBalanceRequest{
		Id:      0,
		JsonRPC: "2.0",
		Method:  "tenderly_addBalance",
		Params:  params,
	})
	if err != nil {
		log.Fatal(err)
	}

	r, err = http.NewRequest("POST", forkRpcUrl, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")

	res, err = httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return forkRpcUrl, forkResponse.SimulationFork.Id
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

func sendTx(forkClient *ethclient.Client, to common.Address, data []byte) {
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

	forkClient.SendTransaction(context.Background(), signedTx)
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
