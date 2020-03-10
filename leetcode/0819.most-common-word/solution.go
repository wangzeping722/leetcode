package main

import "strings"

func main() {
	mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
}

func mostCommonWord(paragraph string, banned []string) string {

	count := make(map[string]int)

	paragraph = strings.ToLower(paragraph)

	words := strings.FieldsFunc(paragraph, Split)
	for _, v := range words {
		if !isInBanned(v, banned) {
			count[v]++
		}
	}

	var str string
	var num int
	for k, v := range count {
		if v > num {
			num = v
			str = k
		}
	}
	return str
}

func isInBanned(word string, banned []string) bool {
	for _, v := range banned {
		if word == v {
			return true
		}
	}
	return false
}

func selfSplit(s rune) bool {
	if s == ' ' || s == '!' || s == '?' || s == '\'' || s == ',' || s == ';' || s == '.' {
		return true
	}
	return false
}


func Split(s rune) bool {
	if s == ' ' || s == '!' || s == '?' || s == '\'' || s == ',' || s == ';' || s == '.' {
		return true
	}
	return false
}
