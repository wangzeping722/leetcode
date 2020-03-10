package _598_range_addition_ii

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m*n
	}

	minW, minL := ops[0][0], ops[0][1]
	for i:=0;i < len(ops); i++ {
		minW = min(minW, ops[i][0])
		minL = min(minL, ops[i][1])
	}

	return minW*minL
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}