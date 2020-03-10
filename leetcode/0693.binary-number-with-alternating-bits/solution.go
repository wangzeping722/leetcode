package _693_binary_number_with_alternating_bits

func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return n & (n+1) == 0
}
