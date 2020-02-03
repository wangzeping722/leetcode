package _453_minimum_moves_to_equal_array_elements

func minMoves(nums []int) int {
	min := nums[0]
	for i:=1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	res := 0
	for _, num := range nums {
		res += num-min
	}
	return res
}
