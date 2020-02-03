package _367_valid_perfect_square

func isPerfectSquare(num int) bool {
	last := 0
	for i := 1; i < num; i++ {
		if i*i == num {
			return true
		} else if i*i > num && last < num {
			return false
		}
		last = i * i
	}
	return true
}

// 牛顿迭代法求平方根
func isPerfectSquare1(num int) bool {
	if num < 2 {
		return true
	}

	i := num / 2
	for i*i > num {
		i = (num/i + i) / 2
	}
	return i*i == num
}
