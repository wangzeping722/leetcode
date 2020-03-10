package _633_sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	for a :=0; a *a <=c; a++ {
		b := int(math.Sqrt(float64(c-a*a)))
		if b*b + a*a == c {
			return true
		}
	}
	return false
}

func judgeSquareSum1(c int) bool {
	for a :=0; a *a <=c; a++ {
		b := a - c*c
		if binarySearch(0, b, b) {
			return true
		}
	}
	return false
}

func binarySearch(left, right, n int) bool {
	if left > right {
		return false
	}

	mid := left + (right - left) / 2
	if mid*mid == n {
		return true
	}
	if mid * mid > n {
		return binarySearch(left, mid-1, n)
	}
	return binarySearch(mid+1, right, n)
}