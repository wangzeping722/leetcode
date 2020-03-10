package _645_set_mismatch

func findErrorNums(nums []int) []int {
	xor := nums[0] ^ 1
	for i := 1; i < len(nums); i++ {
		xor = xor ^ (i + 1) ^ nums[i]
	}
	rightBit := xor & (^(xor-1))
	xor0, xor1 := 0, 0

	for _, num := range nums {
		if num & rightBit != 0 {
			xor1 ^= num
		} else {
			xor0 ^= num
		}
	}
	for i:=1; i <= len(nums); i++ {
		if i & rightBit != 0 {
			xor1 ^= i
		} else {
			xor0 ^= i
		}
	}

	for _, num := range nums {
		if num == xor0 {
			return []int{xor0, xor1}
		}
	}
	return []int{xor1, xor0}
}

func findErrorNums1(nums []int) []int {
	w := make([]int, len(nums))
	ans := []int{0, 0}
	for _, n := range nums {
		if w[n-1] == 0 {
			w[n-1] = 1
		}else {
			ans[0] = n
		}
	}
	for index, n := range w {
		if n == 0 {
			ans[1] = index + 1
		}
	}
	return ans
}