package _500_keyboard_row

import "strings"

func findWords(words []string) []string {
	m := map[byte]int{}
	for _, r := range "QWERTYUIOP" {
		m[byte(r)] = 1 << 1
	}
	for _, r := range "ASDFGHJKL" {
		m[byte(r)] = 1 << 2
	}
	for _, r := range "ZXCVBNM" {
		m[byte(r)] = 1 << 3
	}

	res := []string{}

	for _, word := range words {
		w := strings.ToUpper(word)
		temp := m[w[0]]
		for i:=1; i <len(w); i++ {
			temp &= m[w[i]]
		}
		if temp > 0 {
			res = append(res, word)
		}
	}
	return res
}
