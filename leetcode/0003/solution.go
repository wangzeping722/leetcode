package _003

import "strings"

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	m := make(map[byte]bool)
//	ans, start, end := 0, 0, 0
//	for start < len(s) && end < len(s) {
//		if !m[s[end]] {
//			ans = max(ans, end-start+1)
//			m[s[end]] = true
//			end++
//		} else {
//			delete(m, s[start])
//			start++
//		}
//	}
//	return ans
//}
//
//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}


func lengthOfLongestSubstring(s string) int {
	ans, left, right := 0, 0, 0
	s1 := s[left:right]
	for ; right < len(s); right++ {
		if idx := strings.IndexByte(s1, s[right]); idx != -1 {
			left += idx + 1
		}
		s1 = s[left:right+1]
		if len(s1) > ans {
			ans = len(s1)
		}
	}
	return ans
}