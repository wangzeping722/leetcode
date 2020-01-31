package _292_nim_game


func canWinNim1(n int) bool {
	return n % 4 != 0
}


//递归版本动态规划
func canWinNim2(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 || n == 2 || n == 3 {
		return true
	}
	return !(canWinNim2(n-1) && canWinNim2(n-2) && canWinNim2(n-3))
}


