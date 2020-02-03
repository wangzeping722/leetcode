package _459_repeated_substring_pattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	s1 := s + s
	return strings.Contains(s1[1:len(s1)-1], s)
}
