package problem0070

// 递归，超时
func climbStairs(n int) int {
	if n == 0 || n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs1(n int) int {
	if n == 0 || n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	l1, l2 := 1, 2
	for i := 3; i <= n; i++ {
		l1, l2 = l2, l1+l2
	}

	return l2
}
