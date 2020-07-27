package _392_is_subsequence


// 双指针
func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	if i != len(s) {
		return false
	}
	return true
}

// dp
func isSubsequenceDP(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)

	if sLen > tLen {
		return false
	}
	if sLen == 0 {
		return true
	}
	dp := make([][]bool, sLen+1)
	dp[0] = make([]bool, tLen+1)
	for i := 0; i < tLen; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= sLen; i++ {
		dp[i] = make([]bool, tLen+1)
		for j := 1; j <= tLen; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[sLen][tLen]
}
