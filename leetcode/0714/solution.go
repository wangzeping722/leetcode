package _714

import "math"

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt64
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, temp-prices[i]-fee)
	}
	return dp_i_0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
