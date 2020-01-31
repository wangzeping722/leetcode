package _290_word_pattern

import "strings"

func wordPattern(pattern string, str string) bool {
	ss := strings.Split(str, " ")
	ps := strings.Split(pattern, "")

	if len(ps) != len(ss) {
		return false
	}
	return isMatch(ps, ss) && isMatch(ss, ps)
}

func isMatch(s1, s2 []string) bool {
	size := len(s1)
	m := make(map[string]string, size)

	for i := 0; i < size; i++ {
		if w, ok := m[s1[i]]; ok {
			if w != s2[i] {
				return false
			}
		} else {
			m[s1[i]] = s2[i]
		}
	}
	return true
}
