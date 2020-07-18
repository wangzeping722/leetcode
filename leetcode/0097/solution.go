package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return len(s3) == 0
	}

	if len(s1) == 0 {
		if len(s2) != len(s3) {
			return false
		}
		return s2 == s3
	}

	if len(s2) == 0 {
		if len(s1) == len(s3) {
			return false
		}
		return s1 == s3
	}

	if s1[0] == s2[0] && s1[0] == s3[0] {
		return isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:])
	} else if s1[0] == s3[0] {
		return isInterleave(s1[1:], s2, s3[1:])
	} else if s2[0] == s3[0] {
		return isInterleave(s1, s2[1:], s3)
	}
	return false
}

func main() {
	isInterleave("aabcc", "dbbca", "aadbbcbcac")
}

func isInterleaveDP1(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}

	dp := make([][]bool, s1Len+1)
	for i := 0; i <= s1Len; i++ {
		dp[i] = make([]bool, s2Len+1)
	}

	dp[0][0] = true
	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[s1Len][s2Len]
}

// 使用滚动数组优化空间复杂度
func isInterleaveDP2(s1 string, s2 string, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}

	dp := make([]bool, s2Len+1)
	dp[0] = true
	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || (dp[j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[s2Len]
}

func isInterleaveDP3(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if t != m+n {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 1; i <= m && s1[i-1] == s3[i-1]; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= n && s2[i-1] == s3[i-1]; i++ {
		dp[0][i] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			p := i + j - 1
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[p]) || (dp[i][j-1] && s2[j-1] == s3[p])
		}
	}
	return dp[m][n]
}
