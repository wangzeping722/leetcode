package _266_palindrome_permutation

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int, len(s))

	for _, r := range s {
		m[r]++
	}

	singleCount := 0
	for _, val := range m {
		if val % 2 == 1 {
			singleCount++
		}
	}

	if singleCount == 1 || singleCount == 0 {
		return true
	}
	return false
}
