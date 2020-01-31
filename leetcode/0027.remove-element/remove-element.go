package problem0027

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	last := 0
	for i:=0; i < len(nums); i++ {
		if val != nums[i] {
			nums[last] = nums[i]
			last++
		}
	}
	return last
}
