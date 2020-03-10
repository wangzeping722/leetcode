package _905_sort_array_by_parity

func sortArrayByParity(A []int) []int {
	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 == 0 {
			i++
		} else if A[j]%2 == 1 {
			j--
		} else {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}
	return A
}
