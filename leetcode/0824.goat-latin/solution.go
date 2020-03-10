package _824_goat_latin

import "strings"

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	res := ""
	for idx, word := range words {
		suffix := ""
		for i := 0; i <= idx; i++ {
			suffix += "a"
		}
		if isConsonant(word[0]) {
			newWord := word[1:] + word[0:1] + "ma"
			res += " " + newWord + suffix
		} else {
			res += " " + word + "ma"+suffix
		}
	}
	return res[1:]
}

var marks string = "aeiouAEIOU"
func isConsonant(b byte) bool {
	if strings.Index(marks, string(b)) != -1 {
		return false
	}
	return true
}
