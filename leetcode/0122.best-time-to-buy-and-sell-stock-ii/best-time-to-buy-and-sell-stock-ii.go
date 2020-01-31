package problem0122


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - prices[i-1]
		if tmp > 0 {
			profit += tmp
		}
	}
	return profit
}