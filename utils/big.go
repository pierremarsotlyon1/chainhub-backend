package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

func Quo(n *big.Int, decimals uint64) float64 {
	balance, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(n), big.NewFloat(0).SetInt(math.BigPow(10, int64(decimals)))).Float64()
	return balance
}

func Mul(n int64, decimals uint64) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(n)), math.BigPow(10, int64(decimals)))
}
