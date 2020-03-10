package _747_largest_number_at_least_twice_of_others

import "math"

func dominantIndex(nums []int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	max, second := math.MinInt32, math.MinInt32
	idx := -1

	for k, v := range nums {
		if max < v {
			max, second = v, max
			idx = k
		} else if second < v {
			second = v
		}
	}
	if second*2 <= max {
		return idx
	}
	return -1
}
