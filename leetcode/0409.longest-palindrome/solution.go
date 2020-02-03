package _409_longest_palindrome

// 使用map解决
func longestPalindrome(s string) int {
	m := make(map[byte]int)

	for _, r := range s {
		m[byte(r)]++
	}
	res := 0
	hasSingle := false
	for _, count := range m {
		if count%2 == 0 {
			res += count
		} else {
			if count > 2 {
				res += count - 1
			}
			hasSingle = true
		}
	}

	if hasSingle {
		res++
	}
	return res
}
