package _859_buddy_strings

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	count := 0
	swap := make([]int, 2)
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if count == 2 {
				return false
			}
			swap[count] = i
			count++
		}
	}

	if count == 2 {
		bytes := []byte(A)
		bytes[swap[0]], bytes[swap[1]] = bytes[swap[1]], bytes[swap[0]]
		if string(bytes) == B {
			return true
		}
	}

	if count == 0 {
		m := make(map[byte]int, len(A))
		for _, b := range A {
			m[byte(b)]++
			if m[byte(b)] == 2 {
				return true
			}
		}
	}

	return false
}
