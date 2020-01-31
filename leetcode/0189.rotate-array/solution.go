package problem0189

// 三次旋转法
func rotate(nums []int, k int)  {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for l,r := 0, len(nums)-1; l < r; l, r = l+1,r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}