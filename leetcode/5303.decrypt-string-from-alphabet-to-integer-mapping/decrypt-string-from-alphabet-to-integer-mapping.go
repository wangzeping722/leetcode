package problem5303

import (
	"strconv"
)

func freqAlphabets(s string) string {
	var res []byte
	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			atoi, _ := strconv.Atoi(s[i : i+2])
			res = append(res, 'a'+byte(atoi)-1)
			i += 3
		} else {
			res = append(res, s[i]+'a'-'1')
			i++
		}
	}
	return string(res)
}
