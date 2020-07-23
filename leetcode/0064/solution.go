package main


func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))

	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(grid[0])-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
/*
[[1,3,1],[1,5,1],[4,2,1]]
[[0,1],[1,0]]
 */
func main() {
	minPathSum([][]int{
		[]int{0,1},
		[]int{1,0},
	})
}
