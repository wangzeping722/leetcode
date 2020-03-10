package _123

func maxProfit(prices []int) int {
	n := len(prices)
	maxK := 2
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j <= maxK; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for k := maxK; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[0]
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][maxK][0]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
