package _383_ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	rb := []byte(ransomNote)
	mb := []byte(magazine)
	m := make(map[byte]int, len(mb))

	for _, b := range mb {
		m[b]++
	}

	for _, b := range rb {
		if count, ok := m[b]; !ok || count == 0 {
			return false
		}
		m[b]--
	}
	return true
}
