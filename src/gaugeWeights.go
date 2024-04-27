package src

import (
	"encoding/json"
	"fmt"
	"log"
	"main/contracts/crv"
	"main/contracts/curveGauge"
	"main/contracts/curvePool"
	"main/contracts/erc20"
	"main/interfaces"
	"main/utils"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const GAUGES_PATH = "./data/gauges.json"
const GAUGES_CONFIG_PATH = "./data/configs/gauges-config.json"

func FetchGaugeWeights() {

	fmt.Println("fetch gauges weights")

	gaugesConfig := utils.ReadConfig(GAUGES_CONFIG_PATH)

	mainnetRpcUrl := "https://eth.public-rpc.com"

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
					fmt.Println("error virtual price", poolAddress.Hex())
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
		}
	}

	writeGauges(gauges)

	utils.WriteConfig(gaugesConfig, 0, GAUGES_CONFIG_PATH)
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
