package problem0009

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x <= 9 {
		return true
	}
	t := x
	y := 0
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return t == y
}

func isPalindrome_str(x int) bool{
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i,j := 0, len(s) - 1; i < j; i,j = i+1, j-1{
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
