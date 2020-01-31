package solution

func titleToNumber(s string) int {
	sum := 0
	for _, c := range s {
		sum = sum * 26 + int(c-'A'+1)
	}

	return sum
}
