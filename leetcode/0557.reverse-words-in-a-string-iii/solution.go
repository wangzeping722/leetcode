package _557_reverse_words_in_a_string_iii


// 双指针法

func reverseWords(s string) string {
	bytes := []byte(s)
	for i, j := 0, 0; j < len(s); j++ {
		if s[j] == ' ' || j == len(s)-1{
			if j == len(s) - 1 {
				j++
			}
			for left, right := i, j-1; left < right; left, right = left+1, right-1 {
				bytes[left], bytes[right] = bytes[right], bytes[left]
			}
			i = j+1
		}
	}

	return string(bytes)
}