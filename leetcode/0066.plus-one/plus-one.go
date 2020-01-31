package problem0066

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, 1, len(digits)+1)
	res[0] = 1
	return append(res, digits...)
}
