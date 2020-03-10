package _728_self_dividing_numbers



func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for left <= right {
		num := left
		for num > 0 {
			t := num%10
			if t==0 ||left%t!=0 {
				break
			}
			num /= 10
		}

		if num == 0 {
			res = append(res, left)
		}
		left++
	}
	return res
}