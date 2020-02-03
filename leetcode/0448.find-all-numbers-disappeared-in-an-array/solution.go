package _448_find_all_numbers_disappeared_in_an_array

import "math"

// TODO 不能使用额外空间，首先就要想到在元素组上进行标记
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		num := int(math.Abs(float64(nums[i])))
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	res := []int{}
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
