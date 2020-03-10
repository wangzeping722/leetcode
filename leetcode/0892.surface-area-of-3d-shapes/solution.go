package _892_surface_area_of_3d_shapes

func surfaceArea(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 0 {
				ans += grid[i][j]*4 + 2 // 四个侧面，两个表面
			}
			if i-1 >= 0 && grid[i-1][j] > 0 {
				ans -= min(grid[i][j], grid[i-1][j])
			}
			if i+1 < len(grid) && grid[i+1][j] > 0 {
				ans -= min(grid[i][j], grid[i+1][j])
			}
			if j-1 >= 0 && grid[i][j-1] > 0 {
				ans -= min(grid[i][j], grid[i][j-1])
			}
			if j+1 < len(grid) && grid[i][j+1] > 0 {
				ans -= min(grid[i][j], grid[i][j+1])
			}
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
