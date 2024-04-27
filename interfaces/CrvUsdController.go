package interfaces

import "github.com/ethereum/go-ethereum/common"

type CrvUsdController struct {
	Address                 common.Address `json:"address"`
	Name                    string         `json:"name"`
	CollateralTokenAddress  common.Address `json:"collateralTokenAddress"`
	CollateralTokenDecimals uint8          `json:"collateralTokenDecimals"`
	CrvUsdData              CrvUsdData     `jso:"crvUSDData"`
	TotalFeesCollectedUSD   float64        `json:"totalFeesCollectedUSD"`
}

type CrvUsdData struct {
	ControllerAddress   common.Address `json:"controllerAddress"`
	CurrentAPR          float64        `json:"currentApr"`
	FutureAPR           float64        `json:"futureApr"`
	BlockNumber         uint64         `json:"blocknumber"`
	BlockTimestamp      uint64         `json:"blocktimestamp"`
	TotalDept           float64        `json:"totalDept"`
	CollateralUSD       float64        `json:"collateralUSD"`
	FeesPendingUSD      float64        `json:"feesPendingUSD"`
	FeesCollectedUSD    float64        `json:"feesCollectedUSD"`
	AmmAddress          common.Address `json:"ammAddress"`
	MonetaryAddress     common.Address `json:"monetaryAddress"`
	CrvUSDOracleAddress common.Address `json:"crvUSDOracleAddress"`
}
