package main

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

func main() {
	pivotIndex([]int{-1,-1,-1,-1,-1,0})
}
