package _771_jewels_and_stones


func numJewelsInStones(J string, S string) int {
	m := make(map[byte]bool, len(J))
	for _, b := range J {
		m[byte(b)] = true
	}

	count := 0
	for _, b := range S {
		if m[byte(b)] {
			count++
		}
	}
	return count
}
