package problem0119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{}
	}

	last := []int{1}
	for i := 1; i <= rowIndex; i++ {
		last = append(last, 1)
		for k := i - 1; k > 0; k-- {
			last[k] += last[k-1]
		}
	}
	return last
}
