package problem0190

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	c := 32
	for num > 0 {
		tmp := num & 1
		res = (res << 1) | tmp
		num >>= 1
		c--
	}
	res <<= c
	return res
}
