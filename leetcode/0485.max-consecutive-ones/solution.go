package _485_max_consecutive_ones


func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if res < count {
				res = count
			}
			count = 0
		}
	}
	if res > count {
		return res
	}
	return count
}