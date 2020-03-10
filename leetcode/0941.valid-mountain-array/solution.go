package _941_valid_mountain_array

func validMountainArray(A []int) bool {
	N := len(A)
	i := 0
	for i < N-1 && A[i] < A[i+1] {
		i++
	}
	if i == 0 || i == N-1 {
		return false
	}

	for i < N-1 && A[i] > A[i+1] {
		i++
	}
	if i != N-1 {
		return false
	}

	return true
}
