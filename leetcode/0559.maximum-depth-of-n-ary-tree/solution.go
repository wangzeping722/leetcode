package _559_maximum_depth_of_n_ary_tree

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i:=0; i<len(nums)-1; i = i+2 {
		sum += nums[i]
	}
	return sum
}
