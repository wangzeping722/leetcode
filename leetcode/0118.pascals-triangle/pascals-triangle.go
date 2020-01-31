package problem0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := make([][]int, 0, numRows)
	last := []int{1}
	res = append(res, last)
	for i := 1; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 1; j < i; j++ {
			temp[j] = last[j-1] + last[j]
		}
		res = append(res, temp)
		last = temp
	}
	return res
}
