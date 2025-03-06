package interfaces

import "github.com/ethereum/go-ethereum/common"

type VmV2Config struct {
	ChainId            int
	DefilammaChainName string
	DrpcChainName      string
	Markets            []VmMarket
}

type VmMarket struct {
	Addresses           []common.Address
	BlockNumberDeployed int
}
