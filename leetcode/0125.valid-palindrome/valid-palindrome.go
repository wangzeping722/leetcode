package problem0125

import "strings"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; {
		if !inRange(s[left]) {
			left++
			continue
		}
		if !inRange(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}


func inRange(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
