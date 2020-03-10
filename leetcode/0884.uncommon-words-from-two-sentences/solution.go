package _884_uncommon_words_from_two_sentences

import "strings"

func uncommonFromSentences(A string, B string) []string {
	m := make(map[string]int)
	sa := strings.Split(A, " ")
	sb := strings.Split(B, " ")
	for i:=0; i<len(sa); i++ {
		m[sa[i]]++
	}
	for i:=0; i<len(sb); i++ {
		m[sb[i]]++
	}
	res := []string{}
	for k,v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
