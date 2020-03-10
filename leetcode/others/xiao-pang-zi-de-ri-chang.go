package others

import "sort"

func otaku(nums []int) bool {
	sort.Ints(nums)
	cur := 0
	for cur < len(nums) && nums[cur] == 0 {
		cur++
	}

	for i := cur + 1; i < len(nums); i++ {
		for cur > 0 && nums[i-1]+1 < nums[i] {
			nums[i-1]++
		}
		if nums[i] != nums[i-1]+1 {
			return false
		}
	}
	return true
}
