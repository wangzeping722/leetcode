package _961_n_repeated_element_in_size_2n_array



func repeatedNTimes(A []int) int {
	m := make(map[int]bool)
	for _, num := range A {
		if !m[num] {
			m[num] = true
		} else {
			return num
		}
	}
	return -1
}

func repeatedNTimes1(A []int) int {
	for i:=0; i< len(A)-2; i++ {
		if A[i] == A[i+1] || A[i] == A[i+2] {
			return A[i]
		}
	}
	return -1
}