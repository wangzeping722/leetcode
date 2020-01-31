package problem0069

// 二分法 TODO  多写几次
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	for left < right {
		mid := left + (right-left+1)>>1
		sqrt := mid * mid
		if sqrt == x {
			return mid
		} else if sqrt > x {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

// 牛顿法
func mySqrt1(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res)/2
	}
	return res
}
