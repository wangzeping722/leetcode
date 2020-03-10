package _994_rotting_oranges

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	// 找出所有点
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				bfs(grid, i, j, 2)
			}
		}
	}

	var max = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
	}
	if max == 0 {
		return 0
	}
	return max - 2
}

// 兼顾更新最佳路径
func bfs(grid [][]int, i, j, depth int) {
	grid[i][j] = depth
	if i > 0 && (grid[i-1][j] == 1 || grid[i-1][j]-grid[i][j] > 1) {
		bfs(grid, i-1, j, depth+1)
	}

	if i < len(grid)-1 && (grid[i+1][j] == 1 || grid[i+1][j]-grid[i][j] > 1) {
		bfs(grid, i+1, j, depth+1)
	}

	if j > 0 && (grid[i][j-1] == 1 || grid[i][j-1] - grid[i][j] > 1) {
		bfs(grid, i, j-1, depth+1)
	}

	if j < len(grid[0])-1 && (grid[i][j+1] == 1 || grid[i][j+1]-grid[i][j] > 1) {
		bfs(grid, i, j+1, depth+1)
	}
}
