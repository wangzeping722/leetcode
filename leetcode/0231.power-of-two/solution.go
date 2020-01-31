package problem0231

func isPowerOfTwo(n int) bool {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	if count == 1 {
		return true
	}
	return false
}


// 错位运算更加好
func isPowerOfTwo1(n int) bool {
	if n <= 0 {
		return false
	}

	if n & (n-1) == 0 {
		return true
	}
	return false
}