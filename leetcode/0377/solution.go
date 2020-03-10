package _377

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for k := 0; k < len(nums); k++ {
			if i-nums[k] >= 0 {
				dp[i] += dp[i-nums[k]]
			}
		}
	}
	return dp[target]
}
