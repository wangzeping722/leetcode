package _696_count_binary_substrings

func countBinarySubstrings(s string) int {
	last, cur, res := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cur++
		} else {
			last = cur
			cur = 1
		}
		if last >= cur {
			res++
		}
	}
	return res
}
