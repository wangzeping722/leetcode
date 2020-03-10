package _949_largest_time_for_given_digits

import "fmt"

func largestTimeFromDigits(A []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i != j {
				for k := 0; k < 4; k++ {
					if k != i && k != j {
						l := 6 - i - j - k
						hour := 10*A[i] + A[j]
						min := 10*A[k] + A[l]
						if hour < 24 && min < 60 {
							ans = max(ans, hour*60+min)
						}
					}
				}
			}
		}
	}
	if ans >= 0 {
		return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
	}
	return ""
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
