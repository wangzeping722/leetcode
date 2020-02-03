package _461_hamming_distance

func hammingDistance(x int, y int) int {
	temp := x ^ y
	res := 0
	for i:=0; i<32; i++ {
		if temp & 1 == 1 {
			res++
		}
		temp>>=1
	}
	return res
}
