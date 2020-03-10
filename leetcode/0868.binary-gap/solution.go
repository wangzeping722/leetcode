package _868_binary_gap

func binaryGap(N int) int {
	var (
		maxLength int
		cnt       = 0
		flag      bool
	)
	for N > 0 {
		if flag {
			cnt++
		}
		if N&1 == 1 {
			flag = true
			if maxLength < cnt {
				maxLength = cnt
			}
			cnt = 0
		}

		N >>= 1
	}
	return maxLength
}

