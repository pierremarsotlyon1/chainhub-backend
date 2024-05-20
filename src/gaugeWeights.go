package src

import (
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/crv"
	"main/contracts/curveGC"
	"main/contracts/curveGauge"
	"main/contracts/curvePool"
	"main/contracts/erc20"
	"main/interfaces"
	"main/utils"
	"math"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const GAUGES_PATH = "./data/gauges.json"
const GAUGES_TOP_WINNER_LOSER_PATH = "./data/gauges/top-winner-loser.json"
const GAUGES_CONFIG_PATH = "./data/configs/gauges-config.json"

func FetchGaugeWeights(alchemyMainnetRpc string) {

	fmt.Println("fetch gauges weights")

	gaugesConfig := utils.ReadConfig(GAUGES_CONFIG_PATH)

	mainnetRpcUrl := alchemyMainnetRpc

	client, err := ethclient.Dial(mainnetRpcUrl)
	if err != nil {
		panic(err)
	}

	crvContract, err := crv.NewCrv(utils.CRV, client)
	if err != nil {
		panic(err)
	}

	crvRate, err := crvContract.Rate(nil)
	if err != nil {
		panic(err)
	}
	crvRateF := utils.Quo(crvRate, uint64(18))

	gauges := utils.GetAllGauges()

	now := uint64(time.Now().Unix())
	crvPrice := utils.GetHistoricalPriceTokenPrice(utils.CRV, "ethereum", now)

	for i := 0; i < len(gauges); i++ {
		gauge := gauges[i]

		if gauge.IsKilled || gauge.HasNoCrv {
			continue
		}

		rpcUrl := mainnetRpcUrl
		chainToken := "ethereum"

		if gauge.IsPool {
			if gauge.SideChain {
				index := strings.Index(gauge.Name, "-")
				if index == -1 {
					continue
				}

				chain := gauge.Name[0:index]
				rpcUrl = utils.GetPublicRpcUrl(chain)
				chainToken = utils.GetDefilammaChain(chain)
				if len(rpcUrl) == 0 {
					continue
				}
			}

			client, err := ethclient.Dial(rpcUrl)
			if err != nil {
				fmt.Println("Error when creating rpc ", rpcUrl)
				continue
			}

			gaugeAddress := common.HexToAddress(gauge.Gauge)
			poolAddress := common.HexToAddress(gauge.Pool)

			poolContract, err := curvePool.NewCurvePool(poolAddress, client)
			if err != nil {
				fmt.Println("Error when creating curve pool contract ", poolAddress.Hex())
				continue
			}

			virtualPrice, err := poolContract.VirtualPrice(nil)
			if err != nil {
				virtualPrice, err = poolContract.GetVirtualPrice(nil)
				if err != nil {
					fmt.Println("error virtual price", poolAddress.Hex(), err)
					continue
				}
			}

			virtualPriceF := utils.Quo(virtualPrice, uint64(18))

			lpAddress := common.HexToAddress(gauge.LP)

			totalFees := 0.0

			lpContract, err := curvePool.NewCurvePool(lpAddress, client)
			if err != nil {
				fmt.Println("Error when creating lp contract ", lpAddress.Hex())
				continue
			}

			lpTotalSupply, err := lpContract.TotalSupply(nil)
			if err != nil {
				fmt.Println("Error when fetching lp total supply ", lpAddress.Hex())
				continue
			}

			lpTotalSupplyF := utils.Quo(lpTotalSupply, 18)

			index := int64(0)
			totalDollar := 0.0
			useEndpoint := false

			for {
				coin, err := poolContract.Coins(nil, big.NewInt(index))
				if err != nil {
					break
				}

				// Search decimals token in curve api
				decimals, err := utils.GetTokenDecimals(client, coin)
				if err != nil {
					fmt.Println(coin.Hex())
					continue
				}

				coinPrice := utils.GetHistoricalPriceTokenPrice(coin, chainToken, now)
				if coinPrice == 0 {
					if len(gauge.GaugeCrvApy) == 2 {
						gauges[i].CurrentMinApr = gauge.GaugeCrvApy[0]
						gauges[i].CurrentMaxApr = gauge.GaugeCrvApy[1]
					}

					if len(gauge.GaugeFutureCrvApy) == 2 {
						gauges[i].NextMinApr = gauge.GaugeFutureCrvApy[0]
						gauges[i].NextMaxApr = gauge.GaugeFutureCrvApy[1]
					}
					useEndpoint = true
					break
				}

				coinBalance, err := poolContract.Balances(nil, big.NewInt(index))
				if err != nil {
					fmt.Println("Error when fetching balances ", poolAddress.Hex(), i)
					continue
				}

				balance := utils.Quo(coinBalance, uint64(decimals))

				coinContract, err := erc20.NewErc20(coin, client)
				if err == nil {
					poolCoinBalance, err := coinContract.BalanceOf(nil, poolAddress)
					if err == nil {
						poolCoinBalanceF := utils.Quo(poolCoinBalance, uint64(decimals))
						totalFees += ((poolCoinBalanceF - balance) * coinPrice)
					}
				}

				totalDollar += balance * coinPrice
				index++
			}

			if useEndpoint {
				continue
			}

			lpPrice := totalDollar / lpTotalSupplyF
			if math.IsNaN(lpPrice) {
				lpPrice = 0
			}

			gaugeLpBalanceBN, err := lpContract.BalanceOf(nil, gaugeAddress)
			if err != nil {
				fmt.Println("Error when fetch LP balance in gauge ", gaugeAddress.Hex())
				continue
			}

			gaugeContract, err := curveGauge.NewCurveGauge(gaugeAddress, client)
			if err != nil {
				fmt.Println("Error when creating gauge contract ", gaugeAddress.Hex())
				continue
			}

			workingSupply, err := gaugeContract.WorkingSupply(nil)
			if err != nil {
				fmt.Println("Error when fetching gauge working supply ", gaugeAddress.Hex())
				continue
			}

			workingSupplyF := utils.Quo(workingSupply, uint64(18))

			gw, success := big.NewInt(0).SetString(gauge.GaugeController.GaugeRelativeWeight, 10)
			if !success {
				fmt.Println("Error when convert gauge weight to int")
				continue
			}
			currentGaugeWeightF := utils.Quo(gw, uint64(18))

			gw, success = big.NewInt(0).SetString(gauge.GaugeController.GaugeFutureRelativeWeight, 10)
			if !success {
				fmt.Println("Error when convert new gauge weight to int")
				continue
			}
			nextGaugeWeightF := utils.Quo(gw, uint64(18))

			gauges[i].CurrentMinApr = (crvPrice * crvRateF * currentGaugeWeightF * 12614400) / (workingSupplyF * lpPrice * virtualPriceF) * 100
			gauges[i].CurrentMaxApr = (crvPrice * crvRateF * currentGaugeWeightF * 31536000) / (workingSupplyF * lpPrice * virtualPriceF) * 100
			gauges[i].NextMinApr = (crvPrice * crvRateF * nextGaugeWeightF * 12614400) / (workingSupplyF * lpPrice * virtualPriceF) * 100
			gauges[i].NextMaxApr = (crvPrice * crvRateF * nextGaugeWeightF * 31536000) / (workingSupplyF * lpPrice * virtualPriceF) * 100
			gauges[i].PendingFees = totalFees
			gauges[i].AmountStakedUSD = utils.Quo(gaugeLpBalanceBN, 18) * lpPrice
			gauges[i].LpTokenPrice = lpPrice
			gauges[i].WorkingSupply = workingSupplyF
			gauges[i].VirtualPrice = virtualPriceF
			gauges[i].CanCalculateInTheUI = true

			if utils.IsInfinite(gauges[i].CurrentMinApr) {
				gauges[i].CurrentMinApr = 0
			}
			if utils.IsInfinite(gauges[i].CurrentMaxApr) {
				gauges[i].CurrentMaxApr = 0
			}
			if utils.IsInfinite(gauges[i].NextMinApr) {
				gauges[i].NextMinApr = 0
			}
			if utils.IsInfinite(gauges[i].NextMaxApr) {
				gauges[i].NextMaxApr = 0
			}
		} else {
			if len(gauge.GaugeCrvApy) == 2 {
				gauges[i].CurrentMinApr = gauge.GaugeCrvApy[0]
				gauges[i].CurrentMaxApr = gauge.GaugeCrvApy[1]
			}

			if len(gauge.GaugeFutureCrvApy) == 2 {
				gauges[i].NextMinApr = gauge.GaugeFutureCrvApy[0]
				gauges[i].NextMaxApr = gauge.GaugeFutureCrvApy[1]
			}

			if common.IsHexAddress(gauges[i].Gauge) && common.IsHexAddress(gauges[i].LendingVaultAddress) {
				lendingVaultContract, err := erc20.NewErc20(common.HexToAddress(gauges[i].LendingVaultAddress), client)
				if err == nil {
					gaugeLpBalanceBN, err := lendingVaultContract.BalanceOf(nil, common.HexToAddress(gauges[i].Gauge))
					if err == nil {
						lpPrice := gauges[i].LpTokenPrice
						gauges[i].AmountStakedUSD = utils.Quo(gaugeLpBalanceBN, 18) * lpPrice
					}
				}
			}

			gauges[i].CanCalculateInTheUI = false
		}
	}

	writeGauges(gauges)
	topWinnersLosers(client, gauges)
	utils.WriteConfig(gaugesConfig, 0, GAUGES_CONFIG_PATH)
}

func topWinnersLosers(client *ethclient.Client, gauges []interfaces.Gauge) {
	gcContract, err := curveGC.NewCurveGC(utils.CURVE_GC_ADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}

	totalWeightBN, err := gcContract.GetTotalWeight(nil)
	if err != nil {
		log.Fatal(err)
	}

	totalWeight := utils.Quo(totalWeightBN, 36)

	now := uint64(time.Now().Unix())
	currentPeriod := uint64(now/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC
	nextPeriod := currentPeriod + utils.WEEK_TO_SEC

	allGaugesTopLoser := readMapGaugesTopWinerLoser()

	_, exists := allGaugesTopLoser[currentPeriod]
	if !exists {
		allGaugesTopLoser[currentPeriod] = make([]interfaces.GaugeTopLoser, 0)
	}

	_, exists = allGaugesTopLoser[nextPeriod]
	if !exists {
		allGaugesTopLoser[nextPeriod] = make([]interfaces.GaugeTopLoser, 0)
	}

	for _, gauge := range gauges {
		// Previous period
		gw, success := big.NewInt(0).SetString(gauge.GaugeController.GaugeRelativeWeight, 10)
		if !success {
			fmt.Println("Error when convert gauge weight to int")
			continue
		}
		relativeWeight := utils.Quo(gw, uint64(18))
		weight := relativeWeight * totalWeight

		found := false
		for i := 0; i < len(allGaugesTopLoser[currentPeriod]); i++ {
			if strings.EqualFold(allGaugesTopLoser[currentPeriod][i].Gauge.Hex(), gauge.Gauge) {

				allGaugesTopLoser[currentPeriod][i].CurrentWeight = weight
				found = true
				break
			}
		}

		if !found {
			allGaugesTopLoser[currentPeriod] = append(allGaugesTopLoser[currentPeriod], interfaces.GaugeTopLoser{
				Gauge:         common.HexToAddress(gauge.Gauge),
				CurrentWeight: weight,
			})
		}

		// Next period
		gw, success = big.NewInt(0).SetString(gauge.GaugeController.GaugeFutureRelativeWeight, 10)
		if !success {
			fmt.Println("Error when convert gauge weight to int")
			continue
		}
		relativeWeight = utils.Quo(gw, uint64(18))
		weight = relativeWeight * totalWeight

		found = false
		for i := 0; i < len(allGaugesTopLoser[nextPeriod]); i++ {
			if strings.EqualFold(allGaugesTopLoser[nextPeriod][i].Gauge.Hex(), gauge.Gauge) {

				allGaugesTopLoser[nextPeriod][i].CurrentWeight = weight
				found = true
				break
			}
		}

		if !found {
			allGaugesTopLoser[nextPeriod] = append(allGaugesTopLoser[nextPeriod], interfaces.GaugeTopLoser{
				Gauge:         common.HexToAddress(gauge.Gauge),
				CurrentWeight: weight,
			})
		}
	}

	writeGaugesTopWinerLoser(allGaugesTopLoser)
}

func GetHistoricalTopWinnersLosers(client *ethclient.Client, start uint64) {
	gauges := utils.GetAllGauges()

	now := uint64(time.Now().Unix())
	currentPeriod := uint64(now/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC
	nextCurrentPeriod := currentPeriod + utils.WEEK_TO_SEC

	// 1st january 2024 1704063607
	start = uint64(start/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC

	gcContract, err := curveGC.NewCurveGC(utils.CURVE_GC_ADDRESS, client)
	if err != nil {
		log.Fatal(err)
	}

	allGaugesTopLoser := readMapGaugesTopWinerLoser()

	for {
		if start >= nextCurrentPeriod {
			break
		}

		blockNumber := utils.GetBlockNumberByTimestamp("ethereum", start) + 100
		currentOpts := new(bind.CallOpts)
		currentOpts.BlockNumber = big.NewInt(int64(blockNumber))

		totalWeightBN, err := gcContract.GetTotalWeight(currentOpts)
		if err != nil {
			log.Fatal(err)
		}

		totalWeight := utils.Quo(totalWeightBN, 36)

		gaugesTopLoser := make([]interfaces.GaugeTopLoser, 0)
		for _, gauge := range gauges {

			relativeWeightBN, err := gcContract.GaugeRelativeWeight(currentOpts, common.HexToAddress(gauge.Gauge))
			if err != nil {
				log.Fatal(err)
			}

			relativeWeight := utils.Quo(relativeWeightBN, 18)

			weight := relativeWeight * totalWeight

			gaugesTopLoser = append(gaugesTopLoser, interfaces.GaugeTopLoser{
				Gauge:         common.HexToAddress(gauge.Gauge),
				CurrentWeight: weight,
			})
		}

		currentPeriod := uint64(start/utils.WEEK_TO_SEC) * utils.WEEK_TO_SEC
		allGaugesTopLoser[currentPeriod] = gaugesTopLoser

		start = start + utils.WEEK_TO_SEC
	}

	writeGaugesTopWinerLoser(allGaugesTopLoser)
}

func writeGauges(gauges []interfaces.Gauge) {
	file, err := json.Marshal(gauges)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(GAUGES_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func writeGaugesTopWinerLoser(gauges map[uint64][]interfaces.GaugeTopLoser) {
	file, err := json.Marshal(gauges)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(GAUGES_TOP_WINNER_LOSER_PATH, file, 0644); err != nil {
		log.Fatal(err)
	}
}

func readMapGaugesTopWinerLoser() map[uint64][]interfaces.GaugeTopLoser {
	if !utils.FileExists(GAUGES_TOP_WINNER_LOSER_PATH) {
		return make(map[uint64][]interfaces.GaugeTopLoser)
	}

	file, err := os.ReadFile(GAUGES_TOP_WINNER_LOSER_PATH)
	if err != nil {
		log.Fatal(err)
	}

	gauges := make(map[uint64][]interfaces.GaugeTopLoser)
	if err := json.Unmarshal([]byte(file), &gauges); err != nil {
		log.Fatal(err)
	}

	return gauges
}
