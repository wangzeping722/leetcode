package _482_license_key_formatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	countDashs := strings.Count(s, "-")
	countAN := len(s) - countDashs
	if countAN == 0 {
		return ""
	}

	s = strings.ReplaceAll(s, "-", "")
	s = strings.ToUpper(s)

	res :=make([]byte, (countAN+k-1)/k-1+countAN)
	i,j := len(res), len(s)
	for 0 <= j-k {
		copy(res[i-k:i], s[j-k:j])
		if 0 <= i-k-1 {
			res[i-k-1] = '-'
		}
		i -= k+1
		j -= k
	}
	// 处理头部不完整的字符
	if j > 0 {
		copy(res[:j], s[:j])
	}

	return string(res)
}
