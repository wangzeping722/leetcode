package _506_relative_ranks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	res := make([]string, len(nums))
	as := make(athletes, len(nums))

	for i:= range nums {
		as[i] = athlete{
			score: nums[i],
			index: i,
		}
	}
	sort.Sort(as)
	for i, a := range as {
		switch i {
		case 0:
			res[a.index] = "Gold Medal"
		case 1:
			res[a.index] = "Silver Medal"
		case 2:
			res[a.index] = "Bronze Medal"
		default:
			res[a.index] = strconv.Itoa(i + 1)
		}
	}
	return res
}

type athlete struct {
	score, index int
}

type athletes []athlete

func (a athletes) Len() int {
	return len(a)
}

func (a athletes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a athletes) Less(i ,j int) bool {
	return a[i].score > a[j].score
}