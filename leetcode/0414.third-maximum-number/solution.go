package _414_third_maximum_number

import "math"

func thirdMax(nums []int) int {
	max, second, third := nums[0], math.MinInt64, math.MinInt64
	for i:=1; i < len(nums); i++ {
		if nums[i] == max || nums[i] == second || nums[i] == third {
			continue
		} else if nums[i] > max {
			max, second, third = nums[i], max, second
		} else if nums[i] > second {
			second, third = nums[i], second
		} else if nums[i] > third {
			third = nums[i]
		}
	}
	if third == math.MinInt64 {
		return max
	}
	return third
}