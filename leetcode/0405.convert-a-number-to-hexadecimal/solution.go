package _405_convert_a_number_to_hexadecimal

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hexStr := "0123456789abcdef"
	mask := 0xf
	i := 0
	res := ""
	for i < 8 && num != 0 {
		val := num & mask
		res = string(hexStr[val]) + res
		num >>= 4
		i++
	}
	return res
}

func toHex1(num int) string {
	const (
		hexStr = "0123456789abcdef"
		mask   = 0xf
	)
	var res = make([]byte, 8)
	var nonZeroIndex = 7
	for i := 7; i >= 0; i-- {
		val := num & mask
		res[i] = hexStr[val]
		if val > 0 {
			nonZeroIndex = i
		}
		num >>= 4
	}
	return string(res[nonZeroIndex:])
}
