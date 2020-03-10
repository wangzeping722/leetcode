package _504_base_7

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	res := ""
	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	if flag {
		res = "-" + res
	}
	return res
}
