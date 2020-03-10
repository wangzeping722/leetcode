package _942_di_string_match



func diStringMatch(S string) []int {
	res := []int{}
	left, right := 0, len(S)
	for _, b := range S {
		switch b {
		case 'I':
			res = append(res, left)
			left++
		case 'D':
			res = append(res, right)
			right--
		}
	}
	res = append(res, right)
	return res
}