package _074

func minDistance(word1 string, word2 string) int {
	var dp func(int, int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j+1
		}
		if j == -1 {
			return i+1
		}

		if word1[i] == word2[j] {
			return dp(i-1, j-1) // 当前字符相等，什么都不做
		}
		return min(dp(i-1, j), min(dp(i, j-1), dp(i-1, j-1)))+1
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
