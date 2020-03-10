package _796_rotate_string

import "strings"

//func rotateString(A string, B string) bool {
//	if len(A) != len(B) {
//		return false
//	}
//
//	for i := 0; i <= len(B); i++ {
//		if helper(B, i) == A {
//			return true
//		}
//	}
//	return false
//}
//
//func helper(s string, idx int) string {
//	bytes := []byte(s)
//	rotate(bytes[:idx])
//	rotate(bytes[idx:])
//	rotate(bytes)
//	return string(bytes)
//}
//
//func rotate(bytes []byte) {
//	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
//		bytes[i], bytes[j] = bytes[j], bytes[i]
//	}
//}

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
