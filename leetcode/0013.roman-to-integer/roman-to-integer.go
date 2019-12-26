package problem0013

// 罗马数字中小的数字在大的数字的右边。
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	romans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := romans[s[i]]
		sign := 1
		if tmp < last {
			sign = -1
		}
		res += tmp * sign
		last = tmp
	}
	return res
}
