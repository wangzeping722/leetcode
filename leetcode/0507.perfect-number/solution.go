package _507_perfect_number

import "math"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	n := int(math.Sqrt(float64(num)))
	sum := 1
	for i:=2; i<=n; i++ {
		if num%i == 0 {
			sum = sum + i + num/i
			if num/i == n {
				sum -= n
			}
		}
	}
	return num == sum
}
