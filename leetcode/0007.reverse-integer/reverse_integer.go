package problem0007

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	y := 0
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	y = sign * y
	// 判断边界
	if y > math.MaxInt32 || y < math.MinInt32 {
		y = 0
	}
	return y
}
