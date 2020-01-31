package _256_paint_house

// 动态规划
// 前n个房子的最佳话费为min(m1,m2,m3)
func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	dp := make([]int, 3)
	for i := 0; i < len(costs); i++ {
		minA := min(dp[1], dp[2]) + costs[i][0]
		minB := min(dp[0], dp[2]) + costs[i][1]
		minC := min(dp[0], dp[1]) + costs[i][2]
		dp[0], dp[1], dp[2] = minA, minB, minC
	}
	return min3(dp[0], dp[1], dp[2])
}

func min3(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		min = c
	}
	return min
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}