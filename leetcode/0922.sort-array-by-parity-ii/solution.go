package _922_sort_array_by_parity_ii


func sortArrayByParityII(A []int) []int {
	i, j := 0,1
	for j <= len(A)-1 && i <= len(A)-2 {
		if A[i]%2 == 0 {
			i+=2
		} else if A[j]%2 ==1 {
			j+=2
		} else {
			A[i], A[j] = A[j], A[i]
			i+=2
			j+=2
		}
	}
	return A
}