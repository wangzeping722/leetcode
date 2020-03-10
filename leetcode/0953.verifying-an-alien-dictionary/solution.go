package _953_verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	orderM := make(map[byte]int)

	for i := 0;i < len(order);i++ {
		orderM[order[i]] = i + 1
	}
	for i := 1;i < len(words);i++ {
		if !wordsComp(words[i-1], words[i], orderM) {   // words[i-1] <= words[i]           true
			return false
		}
	}
	return true
}

func wordsComp(words1, words2 string, orderM map[byte]int) bool {
	for i := 0;i < len(words1) && i < len(words2);i++ {
		if words1[i] != words2[i] {
			return orderM[words1[i]] <= orderM[words2[i]]
		}
	}
	return len(words1) <= len(words2)
}

