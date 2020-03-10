package _917_reverse_only_letters

func reverseOnlyLetters(S string) string {
	bytes := []byte(S)
	for i,j := 0, len(bytes)-1; i < j; {
		if bytes[i] < 'A' || (bytes[i] > 'Z' && bytes[i] < 'a') || bytes[i] > 'z' {
			i++
		} else if bytes[j] < 'A' || (bytes[j] > 'Z' && bytes[j] < 'a') || bytes[j] > 'z' {
			j--
		} else {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
	}
	return string(bytes)
}
