package _410

import "math"

func splitArray(nums []int, m int) int {
	n := len(nums)
	// dp[i][j] 表示前 i 个被分割成 j 段所能得到的最大连续子数组和的最小值。
	dp := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m + 1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	// 用来计算 i~j 的和
	for i := 0; i < n; i++ {
		sub[i+1] = nums[i] + sub[i]
	}

	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sub[i]-sub[k]))
			}
		}
	}

	return dp[n][m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func splitArray1(nums []int, m int) int {
	left, right := 0, 0
	for _, n := range nums{
		right += n
		if left < n {
			left = n
		}
	}

	for left < right {
		mid := (right-left)/2 + left
		if check(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(nums []int, x, m int) bool {
	sum, cnt := 0, 1
	for i := 0; i < len(nums); i++ {
		if sum + nums[i] > x {
			cnt++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return cnt <= m
}
