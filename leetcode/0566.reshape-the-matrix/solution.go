package _566_reshape_the_matrix

func matrixReshape(nums [][]int, r int, c int) [][]int {
	width, length := len(nums), len(nums[0])
	if width*length != r*c {
		return nums
	}

	count := 0
	temp := make([]int, 0, c)
	res := make([][]int, 0, r)
	for _, ss := range nums {
		for _, num := range ss {
			count++
			temp = append(temp, num)
			if count == c {
				res = append(res, temp)
				temp = make([]int, 0, c)
				count = 0
			}
		}
	}

	return res
}
