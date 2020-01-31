package problem0121

import "math"

// 记录最小值，
// 获取最大利益
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}

	minPrice, maxProfit := math.MaxInt64, 0
	for i :=0 ; i < length; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
