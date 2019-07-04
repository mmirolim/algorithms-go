package checks

import (
	"math"
)

func Eqf64(a, b float64) bool {
	return math.Abs(a-b) < 1.0E-10
}
