package _441_arranging_coins

import "math"

func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+1.0/4)-1.0/2)
}
