package problem0028

// TODO 0028 思路
// KMP
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	next := getNextValueArray([]byte(needle))
	for i, j := 0, 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNextValueArray(sub []byte) (next []int) {
	next = make([]int, len(sub))
	next[0] = 0
	j := 0
	for i := 1; i < len(sub); i++ {
		for j > 0 && sub[i] != sub[j] {
			j = next[j-1] // 回溯
		}
		if sub[i] == sub[j] {
			j++
		}
		next[i] = j
	}
	return
}

func getNext(b []byte) []int {
	next := make([]int, len(b))
	next[0] = -1
	k := -1
	for i := 1; i < len(b); i++ {
		for k != -1 && b[k+1] != b[i] {
			k = next[k]
		}
		if b[k+1] == b[i] {
			k++
		}
		next[i] = k
	}
}
