package _463_island_perimeter

func islandPerimeter(grid [][]int) int {
	prm := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				if j == 0 || grid[i][j-1] == 0 {
					prm += 1
				}
				if i == 0 || grid[i-1][j] == 0 {
					prm += 1
				}
			}
		}
	}
	return prm*2
}
