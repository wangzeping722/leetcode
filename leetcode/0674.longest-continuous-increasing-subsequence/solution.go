package _674_longest_continuous_increasing_subsequence


func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	count := 1
	for right := 1; right < len(nums); right++ {
		if nums[right] > nums[right-1] {
			count++
			if max < count {
				max = count
			}
		} else {
			count = 1
		}
	}
	return max
}
