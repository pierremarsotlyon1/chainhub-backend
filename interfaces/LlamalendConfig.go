package interfaces

import "github.com/ethereum/go-ethereum/common"

type LlamalendConfig struct {
	FactoryAddress  common.Address
	ChainId         uint64
	MarketName      string
	TimestampDeploy uint64
	BlockDeploy     uint64
	RpcUrl          string
	ConfigPath      string
}
