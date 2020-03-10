package _867_transpose_matrix

func transpose(A [][]int) [][]int {
	if len(A) == len(A[0]) {
		for i := 0; i < len(A); i++ {
			for j := i + 1; j < len(A[0]); j++ {
				A[i][j], A[j][i] = A[j][i], A[i][j]
			}
		}
		return A
	} else {
		B := make([][]int, 0, len(A[0]))
		for i := 0; i < len(A[0]); i++ {
			B = append(B, make([]int, len(A)))
			for j := 0; j < len(A); j++ {
				B[i][j] = A[j][i]
			}
		}
		return B
	}
}

func transpose1(A [][]int) [][]int {
	ROW := len(A)
	if ROW == 0 {
		return nil
	}
	COL := len(A[0])
	B := make([][]int, COL)
	for i := 0; i < COL; i++ {
		B[i] = make([]int, ROW)
	}
	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}