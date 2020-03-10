package _628_maximum_product_of_three_numbers

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	a, b := math.MinInt64, math.MinInt64
	if nums[0] < 0 && nums[1] < 0 {
		a = nums[0]*nums[1]*nums[len(nums)-1]
	}

	b = nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	if a > b {
		return a
	}
	return b
}
