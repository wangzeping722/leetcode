package _79_largest_number

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		as := strconv.Itoa(a)
		bs := strconv.Itoa(b)
		return as+bs > bs+as
	})

	if nums[0] == 0 {
		return "0"
	}
	res := []byte{}
	for _, num := range nums {
		res = append(res, strconv.Itoa(num)...)
	}
	return string(res)
}
