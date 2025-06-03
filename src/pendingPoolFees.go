package src

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"main/contracts/crvUsdControllerFactory"
	"main/contracts/curvePoolUtils"
	"main/contracts/curvePoolWithOwner"
	"main/contracts/erc20"
	"main/contracts/feeCollector"
	"main/contracts/multicall"
	"main/contracts/poolOwnerPendingFees"
	"main/interfaces"
	"main/utils"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	withdrawAdminFeesFunctionSignature = "withdraw_admin_fees()"
	claimAdminFeesFunctionSignature    = "claim_admin_fees()"
	PENDING_POOL_FEES_DATA             = "./data/pendingPoolFees/data.json"
	POOL_WITH_WITHDRAW_ADMIN_FEES_DATA = "./data/pendingPoolFees/pool-with-withdraw-admin-fees.json"
	POOL_WITH_CLAIM_ADMIN_FEES_DATA    = "./data/pendingPoolFees/pool-with-claim-admin-fees.json"
	BURNERS_FOR_TOKEN_DATA             = "./data/pendingPoolFees/burnes-for-tokens.json"
	BURNERS_DATA                       = "./data/pendingPoolFees/burners.json"
	COWSWAP_FEES_PER_TOKEN_PATH        = "./data/pendingPoolFees/cowswap-fees-per-token"
	COWSWAP_FEES_TIMESTAMPS            = "./data/pendingPoolFees/cowswap-fees-timestamps.json"
	NB_ADDRESSES_IN_FUNCTION           = 20
	ETHEREUM_CHAIN                     = "ethereum"
	POLYGON_CHAIN                      = "polygon"
	ARBITRUM_CHAIN                     = "arbitrum"
	OPTIMISM_CHAIN                     = "optimism"
	AVALANCHE_CHAIN                    = "avalanche"
	FANTOM_CHAIN                       = "fantom"
)

var (
	CHAINS_APPROVED       = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN, FANTOM_CHAIN}
	COWSWAP_BURNER_CHAINS = []string{ETHEREUM_CHAIN}
	NODE_TENDERLY         = []string{ETHEREUM_CHAIN, ARBITRUM_CHAIN, POLYGON_CHAIN, OPTIMISM_CHAIN}
	FORK_CHAIN_ID         = map[string]string{
		AVALANCHE_CHAIN: "43114",
		FANTOM_CHAIN:    "250",
	}
	MULTICALL_CHAIN = map[string]common.Address{
		ETHEREUM_CHAIN: utils.MULTICALL_MAINNET,
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
			if err := estimatedWithCowswapBurner(earned, client, currentPendingCrvFees, allPools, chain, COWSWAP_CHAINS[chain]); err != nil {
				return
			}
		} else {
			estimatedWithSimulation(earned, client, allPools, chain)
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

func estimatedWithCowswapBurner(earned map[string]float64, mainnetClient *ethclient.Client, currentPendingCrvFees interfaces.PendingPoolFees, allPools []interfaces.CurvePool, chain string, cowswapChain string) error {

	// Get fee collector address
	feeCollectorAddress, exists := FEE_COLLECTOR_PER_CHAIN_COWSWAP_BURNER[chain]
	if !exists {
		return errors.New("default")
	}

	isMainnet := strings.EqualFold(chain, ETHEREUM_CHAIN)

	client := mainnetClient
	var err error
	if !isMainnet {
		rpcUrl := utils.GetPublicRpcUrl(chain)
		if len(rpcUrl) == 0 {
			return errors.New("default")
		}

		client, err = ethclient.Dial(rpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", rpcUrl)
			return errors.New("default")
		}
	}

	// Get burner
	feeCollectorContract, err := feeCollector.NewFeeCollector(feeCollectorAddress, client)
	if err != nil {
		return err
	}

	burner, err := feeCollectorContract.Burner(nil)
	if err != nil {
		return err
	}

	// Get collect end timestamp
	_, endCollectTimeframe, err := feeCollectorContract.EpochTimeFrame(nil, big.NewInt(2))
	if err != nil {
		return err
	}

	timestamps := readMapTimestamp(COWSWAP_FEES_TIMESTAMPS)
	//lastTimestamp, exists := timestamps[chain]
	//if exists {
	//	if lastTimestamp >= endCollectTimeframe.Uint64() {
	//		return
	//	}
	//}

	// Get forward end timestamp
	_, endForwardTimeframe, err := feeCollectorContract.EpochTimeFrame(nil, big.NewInt(8))
	if err != nil {
		return err
	}

	// Current block
	currentBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// Get the block number we will use to query
	collectBlockTimestamp := currentBlock.Time()
	if collectBlockTimestamp > endCollectTimeframe.Uint64() {
		collectBlockTimestamp = endCollectTimeframe.Uint64()
	}

	collectBlockNumber := utils.GetBlockNumberByTimestamp(chain, collectBlockTimestamp)
	if collectBlockNumber == 0 {
		return err
	}

	forwardBlockNumber := utils.GetBlockNumberByTimestamp(chain, endForwardTimeframe.Uint64()-utils.WEEK_TO_SEC)
	if forwardBlockNumber == 0 {
		return err
	}

	// Fetch crvusd controller addresses
	crvUsdControllerContract, err := crvUsdControllerFactory.NewCrvUsdControllerFactory(CRVUSD_CONTROLLER_FACTORY, client)
	if err != nil {
		return err
	}

	nbCollateral, err := crvUsdControllerContract.NCollaterals(nil)
	if err != nil {
		return err
	}

	controllerAddresses := make([]common.Address, nbCollateral.Int64())
	for i := 0; i < int(nbCollateral.Int64()); i++ {
		controllerAddress, err := crvUsdControllerContract.Controllers(nil, big.NewInt(int64(i)))
		if err != nil {
			return err
		}

		controllerAddresses[i] = controllerAddress
	}

	// Fetch collected fees events
	crvUsdFeesCollected := 0.0

	err = utils.ForEachBlockRange(forwardBlockNumber, currentBlock.NumberU64(), 499, func(start uint64, end uint64) error {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: []common.Address{utils.CRVUSD_ADDRESS},
			Topics: [][]common.Hash{
				{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
				{addressToTopic(utils.CRV_USD_FEE_SPLITTER)},
				{addressToTopic(utils.FEE_COLLECTOR_MAINNET)},
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			return fmt.Errorf("FilterLogs error in estimatedWithCowswapBurner: %w", err)
		}

		crvusdContract, err := erc20.NewErc20(utils.CRVUSD_ADDRESS, client)
		if err != nil {
			return fmt.Errorf("ERC20 contract error: %w", err)
		}

		for _, vLog := range logs {
			event, err := crvusdContract.ParseTransfer(vLog)
			if err != nil {
				fmt.Println("ParseTransfer error:", err)
				continue
			}
			crvUsdFeesCollected += utils.Quo(event.Value, 18)
		}

		return nil
	})

	if err != nil {
		fmt.Println("pagination error in estimatedWithCowswapBurner:", err)
		return err
	}

	// Create tokens array
	coinsMap := make(map[common.Address]bool)
	coins := make([]interfaces.Coin, 0)
	for poolIndex := range allPools {
		for i := range allPools[poolIndex].Coins {
			coinAddress := common.HexToAddress(allPools[poolIndex].Coins[i].Address)

			if strings.EqualFold(coinAddress.Hex(), "0xc425c0877C7B0f8179CDBd636b2eFe3204B24bC4") {
				continue
			}

			_, exists := coinsMap[coinAddress]
			if exists {
				continue
			}

			coinsMap[coinAddress] = true
			coins = append(coins, allPools[poolIndex].Coins[i])
		}
	}

	balanceOfAddresses := []common.Address{burner, feeCollectorAddress}
	if isMainnet {
		// Add pool owner
		balanceOfAddresses = append(balanceOfAddresses, utils.POOL_OWNER_PENDING_FEES_MAINNET)
	}

	// Create balance of calls
	curvePoolUtilsAbi, err := curvePoolUtils.CurvePoolUtilsMetaData.GetAbi()
	if err != nil {
		return err
	}

	multicalls := make([]multicall.Multicall3Call, 0)
	for i := range balanceOfAddresses {

		for coinIndex := range coins {
			// Check balance in fee collector
			b, err := curvePoolUtilsAbi.Pack("balanceOf", common.HexToAddress(coins[coinIndex].Address), balanceOfAddresses[i])
			if err != nil {
				return err
			}

			multicalls = append(multicalls, multicall.Multicall3Call{
				Target:   utils.CURVE_POOL_UTILS,
				CallData: b,
			})
		}
	}

	multicallResponses := utils.Multicall(client, multicalls, MULTICALL_CHAIN[chain], big.NewInt(int64(collectBlockNumber)))
	if len(multicallResponses) != len(multicalls) {
		return err
	}

	totalFees := 0.0
	for _, balanceOfAddress := range balanceOfAddresses {
		for _, coin := range coins {
			multicallResponse := multicallResponses[0]
			multicallResponses = multicallResponses[1:]

			if !multicallResponse.Success {
				continue
			}

			decimals := -1
			_decimals, ok := coin.Decimals.(string)
			if ok {
				d, err := strconv.Atoi(_decimals)
				if err == nil {
					decimals = d
				}
			} else {
				_decimals, ok := coin.Decimals.(int)
				if ok {
					decimals = _decimals
				}
			}

			if decimals == -1 {
				continue
			}

			balance := utils.Quo(big.NewInt(0).SetBytes(multicallResponse.ReturnData), uint64(decimals))

			if balanceOfAddress == feeCollectorAddress && common.HexToAddress(coin.Address) == utils.CRVUSD_ADDRESS {
				balance -= crvUsdFeesCollected
			}

			if coin.UsdPrice == nil {
				coin.UsdPrice = 0.0
			}

			usdPrice, ok := coin.UsdPrice.(float64)
			if !ok {
				usdPrice = 0.0
			}

			amount := balance * usdPrice

			// Not swap by Cowswap if total is under $500
			if amount < 500 {
				continue
			}

			totalFees += amount
		}
	}

	// Check fees still in pools
	curvePoolAbi, err := curvePoolWithOwner.CurvePoolMetaData.GetAbi()
	if err != nil {
		return err
	}

	ownerByte, err := curvePoolAbi.Pack("owner")
	if err != nil {
		return err
	}

	multicalls = make([]multicall.Multicall3Call, 0)
	poolsWithWithdrawAdminFee, poolsWithClaimAdminFee := getPoolsWithWithdrawAdminFee(client, allPools, chain)
	for _, poolAddress := range poolsWithWithdrawAdminFee {
		for _, pool := range allPools {
			if strings.EqualFold(pool.Address, poolAddress.Hex()) {

				multicalls = append(multicalls, multicall.Multicall3Call{
					Target:   poolAddress,
					CallData: ownerByte,
				})

				for i, coin := range pool.Coins {
					coinAddress := common.HexToAddress(coin.Address)

					// Fetch token balance in the pool contract (= internal balance + fees)
					b, err := curvePoolUtilsAbi.Pack("balanceOf", coinAddress, poolAddress)
					if err != nil {
						return err
					}

					multicalls = append(multicalls, multicall.Multicall3Call{
						Target:   utils.CURVE_POOL_UTILS,
						CallData: b,
					})

					// Fetch the internal token balances (the real balance to swap for)
					b, err = curvePoolUtilsAbi.Pack("balances", poolAddress, big.NewInt(int64(i)))
					if err != nil {
						return err
					}

					multicalls = append(multicalls, multicall.Multicall3Call{
						Target:   utils.CURVE_POOL_UTILS,
						CallData: b,
					})
				}

				break
			}
		}
	}

	for _, poolAddress := range poolsWithClaimAdminFee {

		claimableClaimAdminFeesByte, err := curvePoolUtilsAbi.Pack("claimable", poolAddress, feeCollectorAddress)
		if err != nil {
			return err
		}

		multicalls = append(multicalls, multicall.Multicall3Call{
			Target:   utils.CURVE_POOL_UTILS,
			CallData: claimableClaimAdminFeesByte,
		})
	}

	feesInPoolResponses := utils.Multicall(client, multicalls, MULTICALL_CHAIN[chain], big.NewInt(int64(collectBlockNumber)))

	for _, poolAddress := range poolsWithWithdrawAdminFee {

		for _, pool := range allPools {
			if strings.EqualFold(pool.Address, poolAddress.Hex()) {
				ownerRes := feesInPoolResponses[0]
				feesInPoolResponses = feesInPoolResponses[1:]

				// If we don't have the method, everyone can call it
				skipPool := true
				// If we don't have the method, everyone can call it
				if !ownerRes.Success {
					skipPool = false
				}

				if skipPool {
					// If we have the method, need to check if we can call it from the pool owner address
					var owner common.Address
					owner.SetBytes(ownerRes.ReturnData)
					if owner == utils.POOL_OWNER_PENDING_FEES_MAINNET {
						skipPool = false
					}
				}

				for _, coin := range pool.Coins {
					balanceRes := feesInPoolResponses[0]
					selfBalance := feesInPoolResponses[1]
					feesInPoolResponses = feesInPoolResponses[2:]

					if skipPool || !balanceRes.Success || !selfBalance.Success {
						continue
					}

					decimals := -1
					_decimals, ok := coin.Decimals.(string)
					if ok {
						d, err := strconv.Atoi(_decimals)
						if err == nil {
							decimals = d
						}
					} else {
						_decimals, ok := coin.Decimals.(int)
						if ok {
							decimals = _decimals
						}
					}

					if decimals == -1 {
						continue
					}

					poolBalance := utils.Quo(big.NewInt(0).SetBytes(balanceRes.ReturnData), uint64(decimals))

					poolSelfBalance := utils.Quo(big.NewInt(0).SetBytes(selfBalance.ReturnData), uint64(decimals))
					if poolSelfBalance == 0 {
						// Issue in the SC function
						continue
					}
					toWithdraw := poolBalance - poolSelfBalance

					if math.IsInf(toWithdraw, 0) || math.IsNaN(toWithdraw) {
						toWithdraw = 0
					}

					if toWithdraw < 0 {
						continue
					}

					if coin.UsdPrice == nil {
						coin.UsdPrice = 0.0
					}

					usdPrice, ok := coin.UsdPrice.(float64)
					if !ok {
						usdPrice = 0.0
					}

					totalFees += toWithdraw * usdPrice
				}
				break
			}
		}
	}

	for _, poolAddress := range poolsWithClaimAdminFee {

		for _, pool := range allPools {
			if strings.EqualFold(pool.Address, poolAddress.Hex()) {
				feesInPoolResponse := feesInPoolResponses[0]
				feesInPoolResponses = feesInPoolResponses[1:]

				if !feesInPoolResponse.Success {
					break
				}

				totalSupply, _ := big.NewInt(0).SetString(pool.TotalSupply, 10)
				if totalSupply == nil {
					break
				}

				totalSupplyInt := utils.Quo(totalSupply, 18)
				if totalSupplyInt == 0 {
					break
				}

				lpClaimable := utils.Quo(big.NewInt(0).SetBytes(feesInPoolResponse.ReturnData), 18)
				totalFees += (lpClaimable * (pool.UsdTotal / totalSupplyInt))
				break
			}
		}
	}

	timestamps[chain] = collectBlockTimestamp
	writeMapTimestamp(timestamps, COWSWAP_FEES_TIMESTAMPS)
	earned[chain] = totalFees
	earned["crvusd"] = crvUsdFeesCollected

	return nil
}

func estimatedWithSimulation(earned map[string]float64, client *ethclient.Client, allPools []interfaces.CurvePool, chain string) {
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
			return
		}

		clientChain, err = ethclient.Dial(rpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", rpcUrl)
			return
		}
	}

	// Generate calldatas for withdraw_many
	poolsWithWithdrawAdminFee, _ := getPoolsWithWithdrawAdminFee(clientChain, pools, chain)
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
			balance = big.NewInt(0)
		}

		newBalance, success = new(big.Int).SetString(lastResult.Trace[0].DecodedOutput[0].Value.(string), 10)
		if !success {
			newBalance = big.NewInt(0)
		}
	} else {
		// Use fork
		forkRpcUrl, forkId, _, err := createFork(FORK_CHAIN_ID[chain], 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer deleteFork(forkId)

		forkClient, err := ethclient.Dial(forkRpcUrl)
		if err != nil {
			fmt.Println("Error when creating rpc ", forkRpcUrl)
			return
		}

		// Simulation de l'exécution de la transaction
		targetTokenContract, err := erc20.NewErc20(TARGET_TOKEN[chain], forkClient)
		if err != nil {
			log.Fatal(err)
		}

		balance, err = targetTokenContract.BalanceOf(nil, FEE_RECEIVERS[chain])
		if err != nil {
			fmt.Println(err)
			balance = big.NewInt(0)
		}

		to := POOL_OWNERS[chain]

		for _, call := range calldatasWithdrawFees {
			sendTx(forkClient, to, call)
		}

		for _, call := range calldatasBurnMany {
			sendTx(forkClient, to, call)
		}

		newBalance, err = targetTokenContract.BalanceOf(nil, FEE_RECEIVERS[chain])
		if err != nil {
			fmt.Println(err)
			newBalance = big.NewInt(0)
		}
	}

	earned[chain] = utils.Quo(new(big.Int).Sub(newBalance, balance), TARKET_TOKEN_DECIMALS[chain])
}

func fetchLastDistribution(client *ethclient.Client, currentBlock uint64, lastBlock uint64, lastTimestamp uint64, lastDistributionAmount uint64) (uint64, uint64) {
	from := lastBlock
	if from == 0 {
		from = 19526061
	}

	timestamp := lastTimestamp

	err := utils.ForEachBlockRange(from, currentBlock, 499, func(start uint64, end uint64) error {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(start)),
			ToBlock:   big.NewInt(int64(end)),
			Addresses: []common.Address{utils.CRVUSD_ADDRESS},
			Topics:    [][]common.Hash{{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")}},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil {
			return fmt.Errorf("FilterLogs error: %w", err)
		}

		for _, vLog := range logs {
			contract, err := erc20.NewErc20(vLog.Address, client)
			if err != nil {
				fmt.Println("ERC20 init error:", err)
				continue
			}

			event, err := contract.ParseTransfer(vLog)
			if err != nil {
				fmt.Println("ParseTransfer error:", err)
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
				fmt.Println("BlockByNumber error:", err)
				continue
			}

			if block.Time()-timestamp > 2*utils.DAY_TO_SEC {
				lastDistributionAmount = 0
			}

			timestamp = block.Time()
			lastDistributionAmount += uint64(utils.Quo(event.Value, 18))
		}

		return nil
	})

	if err != nil {
		fmt.Println("fetchLastDistribution error:", err)
	}

	return timestamp, lastDistributionAmount
}

func getPoolsWithWithdrawAdminFee(client *ethclient.Client, pools []interfaces.CurvePool, chain string) ([]common.Address, []common.Address) {

	poolWithWithdrawAdminFee := readMapBool(POOL_WITH_WITHDRAW_ADMIN_FEES_DATA)
	poolWithClaimAdminFee := readMapBool(POOL_WITH_CLAIM_ADMIN_FEES_DATA)

	// Function withdrawAdminFeesSelector
	withdrawAdminFeesSelector := crypto.Keccak256Hash([]byte(withdrawAdminFeesFunctionSignature)).Hex()[2:10]
	claimAdminFeesSelector := crypto.Keccak256Hash([]byte(claimAdminFeesFunctionSignature)).Hex()[2:10]

	for _, pool := range pools {

		// Withdraw
		_, exists := poolWithWithdrawAdminFee[chain]
		if !exists {
			poolWithWithdrawAdminFee[chain] = make(map[common.Address]bool)
		}

		poolAddress := common.HexToAddress(pool.Address)
		haveIt, exists := poolWithWithdrawAdminFee[chain][poolAddress]
		abi := ""

		if !exists {
			if strings.EqualFold(chain, "ethereum") {
				if len(abi) == 0 {
					_abi, err := utils.GetEtherscanAbi(poolAddress.Hex())
					if err == nil {
						abi = _abi
					}
				}
				haveIt, _ = utils.ContainsWithdrawAdminFees(abi)
				time.Sleep(1 * time.Second)
				if haveIt {
					poolWithWithdrawAdminFee[chain][poolAddress] = true
				}
			}

			if !haveIt {
				bytecode, err := client.CodeAt(context.Background(), poolAddress, nil) // nil pour le dernier bloc connu
				if err != nil {
					fmt.Println(err)
				} else {
					poolWithWithdrawAdminFee[chain][poolAddress] = strings.Contains(common.Bytes2Hex(bytecode), withdrawAdminFeesSelector)
				}
			}
		}

		// Claim
		_, exists = poolWithClaimAdminFee[chain]
		if !exists {
			poolWithClaimAdminFee[chain] = make(map[common.Address]bool)
		}

		haveIt, exists = poolWithClaimAdminFee[chain][poolAddress]
		if !exists {
			if strings.EqualFold(chain, "ethereum") {
				if len(abi) == 0 {
					_abi, err := utils.GetEtherscanAbi(poolAddress.Hex())
					if err == nil {
						abi = _abi
					}
				}
				haveIt, _ = utils.ContainsClaimAdminFees(abi)
				time.Sleep(1 * time.Second)
				if haveIt {
					poolWithClaimAdminFee[chain][poolAddress] = true
				}
			}

			if !haveIt {
				bytecode, err := client.CodeAt(context.Background(), poolAddress, nil) // nil pour le dernier bloc connu
				if err != nil {
					fmt.Println(err)
				} else {
					poolWithClaimAdminFee[chain][poolAddress] = strings.Contains(common.Bytes2Hex(bytecode), claimAdminFeesSelector)
				}
			}
		}
	}

	writeMapBol(poolWithWithdrawAdminFee, POOL_WITH_WITHDRAW_ADMIN_FEES_DATA)
	writeMapBol(poolWithClaimAdminFee, POOL_WITH_CLAIM_ADMIN_FEES_DATA)

	poolsWithWithdrawAdminFee := make([]common.Address, 0)
	poolsWithClaimAdminFee := make([]common.Address, 0)
	for poolAddress, value := range poolWithWithdrawAdminFee[chain] {
		if value {
			poolsWithWithdrawAdminFee = append(poolsWithWithdrawAdminFee, poolAddress)
		}
	}
	for poolAddress, value := range poolWithClaimAdminFee[chain] {
		if value {
			poolsWithClaimAdminFee = append(poolsWithClaimAdminFee, poolAddress)
		}
	}
	return poolsWithWithdrawAdminFee, poolsWithClaimAdminFee
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
	pools := make(map[string]map[common.Address]bool)
	b, err := utils.ReadBucketFile(path[2:])
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &pools); err != nil {
			log.Fatal(err)
		}

		return pools
	}

	if !utils.FileExists(path) {
		return make(map[string]map[common.Address]bool)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

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

	if err := utils.WriteBucketFile(path[2:], pools); err != nil {
		fmt.Println(err)
	}
}

func readMap(path string) map[common.Address]bool {
	pools := make(map[common.Address]bool)
	b, err := utils.ReadBucketFile(path[2:])
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &pools); err != nil {
			log.Fatal(err)
		}

		return pools
	}

	if !utils.FileExists(path) {
		return make(map[common.Address]bool)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

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

	if err := utils.WriteBucketFile(path[2:], pools); err != nil {
		fmt.Println(err)
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

	if err := utils.WriteBucketFile(path[2:], pools); err != nil {
		fmt.Println(err)
	}
}

func readPending3CRVFees() interfaces.PendingPoolFees {
	pendingPoolFees := new(interfaces.PendingPoolFees)
	b, err := utils.ReadBucketFile(PENDING_POOL_FEES_DATA[2:])
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &pendingPoolFees); err != nil {
			log.Fatal(err)
		}

		return *pendingPoolFees
	}

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

	if err := utils.WriteBucketFile(PENDING_POOL_FEES_DATA[2:], data); err != nil {
		fmt.Println(err)
	}
}

func writeMapTimestamp(timestamps map[string]uint64, path string) {
	file, err := json.Marshal(timestamps)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, file, 0644); err != nil {
		log.Fatal(err)
	}

	if err := utils.WriteBucketFile(path[2:], timestamps); err != nil {
		fmt.Println(err)
	}
}

func readMapTimestamp(path string) map[string]uint64 {

	timestamps := make(map[string]uint64)
	b, err := utils.ReadBucketFile(path[2:])
	if err == nil && len(b) > 0 {
		if err := json.Unmarshal(b, &timestamps); err != nil {
			log.Fatal(err)
		}

		return timestamps
	}

	if !utils.FileExists(path) {
		return make(map[string]uint64)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(file), &timestamps); err != nil {
		log.Fatal(err)
	}

	return timestamps
}

func addressToTopic(addr common.Address) common.Hash {
	return common.BytesToHash(addr.Bytes())
}
