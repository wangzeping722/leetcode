package _268_missing_number

import "sort"

//
func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := range nums {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}

func missingNumber1(nums []int) int {
	miss := len(nums)
	for i := range nums {
		miss ^= i ^ nums[i]
	}

	return miss
}