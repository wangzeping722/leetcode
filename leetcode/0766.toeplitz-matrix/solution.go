package _766_toeplitz_matrix


func isToeplitzMatrix(matrix [][]int) bool {
	last := matrix[0]
	for i:=1; i< len(matrix);i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != last[j-1] {
				return false
			}
		}
		last = matrix[i]
	}
	return true
}