package _762_prime_number_of_set_bits_in_binary_representation



func countOne(n int) uint {
	res := 0
	for n != 0 {
		res++
		n &= n-1
	}
	return uint(res)
}

func countPrimeSetBits(L int, R int) int {
	m := 665772 // 0b10100010100010101100
	res := 0
	for i := L; i <= R; i++ {
		res += (m >> countOne(i)) & 1
	}
	return res
}

