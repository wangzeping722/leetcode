package _896_monotonic_array

func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	asc, desc := false, false
	for i := 1; i < len(A); i++ {

		if A[i-1] < A[i] {
			asc = true
		}
		if A[i-1] > A[i] {
			desc = true
		}
		if asc && desc {
			return false
		}
	}

	return true
}
