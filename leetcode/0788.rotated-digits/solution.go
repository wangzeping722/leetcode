package _788_rotated_digits

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if isGoodNum(i) {
			count++
		}
	}
	return count
}

func isGoodNum(n int) bool {
	changed := false
	for n > 0 {
		t := n % 10
		n /= 10
		switch t {
		case 2, 5, 6, 9:
			changed = true
		case 0, 1, 8:
		default:
			return false
		}
	}
	if changed {
		return true
	}
	return false
}
