package _581_shortest_unsorted_continuous_subarray

import "sort"

func findUnsortedSubarray(nums []int) int {
	temps := make([]int, len(nums))
	copy(temps, nums)
	sort.Ints(temps)
	left, right := 0, len(nums)-1

	for left < len(nums) && nums[left] == temps[left] {
		left++
	}
	if left == len(nums) {
		return 0
	}

	for nums[right] == temps[right] {
		right--
	}
	return right - left
}

// 双指针
func findUnsortedSubarray1(nums []int) int {
	length := len(nums)
	min, max := nums[length-1], nums[0]
	l,r := 0,-1

	for i, num := range nums {
		if max > num {
			r = i
		} else {
			max = num
		}

		if min < nums[length-i-1] {
			l = length-i-1
		} else {
			min = nums[length-i-1]
		}
	}
	return r-l+1
}