package _665_non_decreasing_array

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	count := 0
	for i:=0; i< len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}
