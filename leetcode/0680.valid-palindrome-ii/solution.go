package _680_valid_palindrome_ii

func validPalindrome(s string) bool {
	return valid(s, true)

}

func valid(s string, canTrim bool) bool {
	if s == "" {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			break
		}
		left++
		right--
	}

	if left >= right {
		return true
	}

	if canTrim {
		return valid(s[left+1:right+1], false) || valid(s[left:right], false)
	}
	return false
}
