package _300


func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 0)
	dp = append(dp, nums[0])
	for i:=1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			l,r := 0, len(dp)-1
			for l < r {
				mid := (r+l)/2
				if dp[mid] < nums[i] {
					l = mid+1
				} else if dp[mid] > nums[i] {
					r = mid
				} else if dp[mid] == nums[i] {
					l = mid+1
				}
			}
			dp[l] = nums[i]
		}

	}

	return len(dp)
}