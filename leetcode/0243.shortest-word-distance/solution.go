package _243_shortest_word_distance

import "math"

func shortestDistance(words []string, word1 string, word2 string) int {
	minLength := len(words)

	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			for j, word := range words {
				if word == word2 {
					if int(math.Abs(float64(j-i))) < minLength {
						minLength = int(math.Abs(float64(j - i)))
					}
				}
			}
		}
	}
	return minLength
}


func shortestDistance1(words []string, word1 string, word2 string) int {
	i1, i2 := -1, -1

	minLength := len(words)
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			i1 = i
		} else if words[i] == word2 {
			i2 = i
		}

		if i1 != -1 && i2 != -1 {
			if int(math.Abs(float64(i1- i2))) < minLength {
				minLength = int(math.Abs(float64(i1 - i2)))
			}
		}
	}
	return minLength
}