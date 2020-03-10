package lcof

//func findRepeatNumber(nums []int) int {
//	m := make(map[int]bool, 0)
//	for _, num := range nums {
//		if m[num] {
//			return num
//		}
//		m[num] = true
//	}
//	return -1
//}


func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	for i:=0; i<len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]]	= nums[nums[i]], nums[i]
		}
	}
	return 0
}