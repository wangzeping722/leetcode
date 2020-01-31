package _258_add_digits

func addDigits(num int) int {
	for num >= 10 {
		temp := num
		num = 0
		for temp > 0 {
			num += temp % 10
			temp /= 10
		}
	}

	return num
}

func addDigits1(num int) int {
	if num > 9 {
		num = num % 9
		if num == 0 {
			return 9
		}
	}
	return num
}