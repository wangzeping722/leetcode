package _342_power_of_four

func isPowerOfFour(n int) bool {
	for n%4 == 0 {
		n /= 4
	}

	return n == 1
}


func isPowerOfFour1(n int) bool {
	if n <= 0 {
		return false
	}

	if n & (n-1) != 0 {
		return false
	}

	return n & 0x55555555 == n
}