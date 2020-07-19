package _312

// 回溯法
// (n)! 超时
// a 不出来
func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for sl := 1; sl <= n; sl++ {
		for i := 1; i <= n-sl+1; i++ {
			j := i + sl - 1
			for k := i; k <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k-1]+dp[k+1][j]+nums[k]*nums[i-1]*nums[j+1])
			}
		}
	}
	return dp[1][n]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//状态转移方程
// dp[i][j] 从 i 到 j 个气球能够获取的最大硬币数量
// dp[i][j] = dp[i][k-1] + dp[k+1][j] + nums[k] * nums[i-1] * nums[j+1]
