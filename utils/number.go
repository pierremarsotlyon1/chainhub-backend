package utils

import "math"

func IsInfinite(num float64) bool {
	return math.IsInf(num, 0) || math.IsNaN(num)
}
