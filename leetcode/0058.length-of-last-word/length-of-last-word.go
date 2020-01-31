package problem0058


func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {}
	lastIndex := i
	for ; i >= 0 && s[i] != ' '; i-- {}
	return lastIndex - i
}