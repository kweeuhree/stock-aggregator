package utils

import "math"

func RoundFloatToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
