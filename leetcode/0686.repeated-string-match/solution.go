package _686_repeated_string_match

import "strings"

func repeatedStringMatch(A string, B string) int {
	endLen := len(A)*2 + len(B)
	str := A
	for i:=1; len(str) <= endLen; i, str = i+1, str+A{
		if strings.Contains(str, B) {
			return i
		}
	}
	return -1
}
