package problem0035

// 二分查找
func searchInsert(nums []int, target int) int {
	var mid int
	low, high := 0, len(nums)-1
	for low <= high {
		mid = low + (high-low) >> 1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid -1
		} else {
			return mid
		}
	}
	return low
}

