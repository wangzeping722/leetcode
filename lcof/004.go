package lcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	top, right := 0, len(matrix[0])-1
	for top < len(matrix) && right >= 0 {
		if matrix[top][right] == target {
			return true
		} else if matrix[top][right] > target {
			right--
		} else {
			top++
		}
	}
	return false
}
