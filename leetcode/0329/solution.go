package _329

var (
	dirs          = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, columns int
)

// 使用记忆化深度优先搜索
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := range memo {
		memo[i] = make([]int, columns)
	}

	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, memo, i, j))
		}
	}
	return ans
}

func dfs(matrix, memo [][]int, row, column int) int {
	if memo[row][column] != 0 {
		return memo[row][column]
	}

	memo[row][column]++
	// 深度优先遍历左右上下四个方向
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			memo[row][column] = max(memo[row][column], dfs(matrix, memo, newRow, newColumn)+1)
		}
	}
	return memo[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
