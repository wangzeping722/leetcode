package main


func backspaceCompare(S string, T string) bool {
	s := make([]byte, len(S))
	t := make([]byte, len(T))
	left := 0
	for i:=0; i<len(S);i++ {
		if S[i] == '#' {
			if left != 0 {
				left--
			}
		} else {
			s[left] = S[i]
			left++
		}
	}
	ss := string(s[:left])
	left = 0
	for i:=0; i<len(T);i++ {
		if T[i] == '#' {
			if left != 0 {
				left--
			}
		} else {
			t[left] = T[i]
			left++
		}
	}
	tt := string(t[:left])
	return ss == tt
}

func main() {
	backspaceCompare("ab#c", "ad#c")
}