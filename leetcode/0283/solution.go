package _283

func moveZeroes(nums []int) {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[j] = nums[i]
			}
			j++
		}
	}

	for ; j < i; j++ {
		nums[j] = 0
	}
}
