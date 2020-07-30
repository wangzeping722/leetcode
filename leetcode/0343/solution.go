package _343

func integerBreak(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j * dp[i-j], j*(i-j)))
		}
	}

	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}