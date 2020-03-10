package _830_positions_of_large_groups

func largeGroupPositions(S string) [][]int {
	left := 0
	res := [][]int{}
	for i := 1; i < len(S); i++ {
		if S[i] != S[i-1] {
			if i-left >= 3 {
				res = append(res, []int{left, i - 1})
			}
			left = i
		} else if i == len(S)-1 {
			if i-left+1 >= 3 {
				res = append(res, []int{left, i})
			}
		}
	}
	return res
}
