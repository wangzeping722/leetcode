package _733_flood_fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	helper(image, sr, sc, newColor, image[sr][sc])
	return image
}

func helper(image [][]int, sr int, sc int, newColor int, originColor int) {
	if sr == -1 || sr == len(image) ||
		sc == -1 || sc == len(image[0]) || image[sr][sc] != originColor {
		return
	}

	image[sr][sc] = newColor
	helper(image, sr+1, sc, newColor, originColor)
	helper(image, sr-1, sc, newColor, originColor)
	helper(image, sr, sc+1, newColor, originColor)
	helper(image, sr, sc-1, newColor, originColor)
}
