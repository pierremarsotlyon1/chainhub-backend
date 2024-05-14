package interfaces

import "github.com/ethereum/go-ethereum/common"

type Wrapper struct {
	Name        string
	Address     common.Address
	LogoUrl     string
	PoolAddress common.Address
}
