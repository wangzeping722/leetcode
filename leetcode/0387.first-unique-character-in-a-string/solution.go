package _387_first_unique_character_in_a_string

import "strings"

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i:=0; i < len(s); i++ {
		m[s[i]]++
	}
	min := -1
	for k,v := range m {
		if v == 1 && (strings.IndexByte(s, k) < min || min == -1){
			min = strings.IndexByte(s, k)
		}
	}

	return min
}
