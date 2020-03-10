package _746_min_cost_climbing_stairs

// 自己写的dp就是香啊
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		if cost[i-1] + dp[i-1] < cost[i-2] + dp[i-2] {
			dp[i] = cost[i-1] + dp[i-1]
		} else {
			dp[i] = cost[i-2] + dp[i-2]
		}
	}
	return dp[len(cost)]
}

// 超时递归
//func minCostClimbingStairs(cost []int) int {
//	s1 := helper(cost, 0)
//	s2 := helper(cost, 1)
//	if s1 < s2 {
//		return s1
//	}
//	return s2
//}
//
//func helper(cost []int, cur int) int {
//	if cur >= len(cost) {
//		return 0
//	}
//
//	s1 := helper(cost, cur+1) + cost[cur]
//	s2 := helper(cost, cur+2) + cost[cur]
//	if s1 > s2 {
//		return s2
//	}
//	return s1
//}
