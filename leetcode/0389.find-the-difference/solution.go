package _389_find_the_difference

func findTheDifference(s string, t string) byte {
	var res rune

	for _, b := range s {
		res ^= b
	}

	for _, b := range t {
		res ^= b
	}

	return byte(res)
}


func findTheDifference1(s string, t string) byte {
	a := 0
	l := len(s)
	for i := 0; i < l; i++ {
		a = a - int(s[i]) + int(t[i])
	}
	a += int(t[l])
	return byte(a)
}