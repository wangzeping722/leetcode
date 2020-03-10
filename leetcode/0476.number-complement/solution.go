package _476_number_complement

func findComplement(num int) int {
	temp := num
	mask := 1
	for temp > 0 {
		temp >>= 1
		mask <<= 1
	}
	return num ^ (mask-1)
}
