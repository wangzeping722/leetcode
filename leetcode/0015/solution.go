package _015

import "sort"

// 超时
//func threeSum(nums []int) [][]int {
//	res := [][]int{}
//	for i:=0; i< len(nums); i++ {
//		if i != 0 && nums[i] == nums[i-1] {	 // 解决超时
//			continue
//		}
//		m := make(map[int]bool, 0)
//		target1 := -nums[i]
//		for j := 0; j < len(nums); j++ {
//			if i == j {
//				continue
//			}
//			if m[target1-nums[j]] {
//				res = distinctAppend(res, []int{nums[i], nums[j], target1-nums[j]})
//			} else {
//				m[nums[j]] = true
//			}
//		}
//	}
//
//	return res
//}
//
func distinctAppend(res [][]int, arr []int) [][]int {
	sort.Ints(arr)
	for i:=0; i < len(res); i++ {
		if res[i][0] == arr[0] && res[i][1] == arr[1] && res[i][2] == arr[2] {
			return res
		}
	}
	return append(res, arr)
}

func threeSum(nums []int) [][]int {
	if len(nums) <3 {
		return nil
	}

	res := make([][]int, 0)
	sort.Ints(nums)
	for i:=0; i < len(nums); i++ {
		if nums[i] > 0 { 	// 直接退出
			break
		}

		if i > 0 && nums[i] == nums[i-1] {	// 去重
			continue
		}

		l, r := i+1, len(nums)-1
		for l  < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = distinctAppend(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {l++}
				for l < r && nums[r] == nums[r-1] {r--}
				l++
				r--
			} else if sum > 0 {
				r--
			}else if sum < 0 {
				l++
			}

		}
	}
	return res
}

