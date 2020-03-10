package _520_detect_capital

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}
	first, second := word[0], word[1]
	if first >= 'a' && first <= 'z' {
		for _, r := range word {
			if r < 'a' || r > 'z' {
				return false
			}
		}
	} else if second >= 'A' && second <= 'Z' {
		for _, r := range word {
			if r < 'A' || r > 'Z' {
				return false
			}
		}
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
	}

	return true
}
