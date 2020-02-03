package _345_fan_zhuan_zi_fu_chuan_zhong_de_yuan_yin_zi_mu_by_q



func reverseVowels(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; {
		if !isVowel(bytes[i]) {
			i++
			continue
		}

		if !isVowel(bytes[j]) {
			j--
			continue
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}