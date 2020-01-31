package problem0053

// dp:
// 空间复杂度 O(n)
// 时间复杂度 O(n)
func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

// TODO 二分解决
//
//func maxSubArray1(nums []int) int {
//
//}