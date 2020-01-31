package problem0198


// f(k) = Max[f(k-2)+Ak, f(k-1)]
// TODO 动态规划
func rob(nums []int) int {
	r1 := 0
	r2 := 0
	for _, a := range nums {
		temp := r2
		r2 = max(r1+a, r2)
		r1 = temp
	}
	return max(r1, r2)
}

func max(a1, a2 int) int {
	if a1 > a2 {
		return a1
	}
	return a2
}
