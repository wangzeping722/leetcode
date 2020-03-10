package _309

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt64
	dp_pre_0 := 0 // 代表dp[i-2][0],两天前未持有
	for i:=0; i< n; i++ {
		temp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = temp
	}
	return dp_i_0
}


func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}