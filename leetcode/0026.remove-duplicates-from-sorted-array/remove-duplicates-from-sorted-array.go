package problem0026

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	last := 0
	for i := 1; i < len(nums); i++ {
		if nums[last] != nums[i] {
			last++
			nums[last] = nums[i]
		}
	}

	return last + 1
}
