package _709_to_lower_case


func toLowerCase(str string) string {
	res := []byte(str)
	step := 'A'-'a'
	for i:=0; i < len(res);i++ {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] -= byte(step)
		}
	}
	return string(res)
}
