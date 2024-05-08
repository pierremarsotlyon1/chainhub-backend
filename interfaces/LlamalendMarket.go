package interfaces

import (
	"github.com/ethereum/go-ethereum/common"
)

type LlamalendMarket struct {
	Id                     int            `json:"id"`
	ControllerAddress      common.Address `json:"controllerAddress"`
	FactoryAddress         common.Address `json:"factoryAddress"`
	MarketName             string         `json:"marketName"`
	Timestamp              uint64         `json:"ts"`
	TotalDebt              float64        `json:"totalDebt"`
	TotalDebtUSD           float64        `json:"totalDebtUSD"`
	Supplied               float64        `json:"supplied"`
	SuppliedUSD            float64        `json:"suppliedUSD"`
	Collateral             float64        `json:"collateral"`
	CollateralUSD          float64        `json:"collateralUSD"`
	BorrowedTokenAddress   common.Address `json:"borrowedTokenAddress"`
	CollateralTokenAddress common.Address `json:"collateralTokenAddress"`
	ChainId                uint64         `json:"chainId"`
	Loans                  int64          `json:"loans"`
	BorrowApr              float64        `json:"borrowApr"`
	SupplyApr              float64        `json:"supplyApr"`
	HardLiquidation        uint64         `json:"hardLiquidation"`
}
