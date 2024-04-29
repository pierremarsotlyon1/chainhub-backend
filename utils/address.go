package utils

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func IsNullAddress(addr common.Address) bool {
	return strings.EqualFold(addr.Hex(), NULL_ADDRESS.Hex())
}
