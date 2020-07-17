package problem0035

// 二分查找
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left+(right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid+1
		}
	}
	return left
}

