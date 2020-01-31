package problem0014

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	min := len(strs[0])
	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}

	prefix := make([]byte, 0, min)
	for i := 0; i < min; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if c != strs[j][i] {
				return string(prefix)
			}
		}
		prefix = append(prefix, c)
	}
	return string(prefix)
}


func longestCommonPrefix1(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j ++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}