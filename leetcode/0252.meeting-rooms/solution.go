package _252_meeting_rooms

import (
	"sort"
)

// 按开始时间排序
func canAttendMeetings(intervals [][]int) bool {
	for i := range intervals {
		min := i
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] < intervals[min][0] {
				min = j
			}
		}
		intervals[i], intervals[min] = intervals[min], intervals[i]
	}


	for i := 0; i < len(intervals) - 1; i++ {
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}
	return true
}


func canAttendMeetings1(intervals [][]int) bool {
	var d = conf{}
	d.data = intervals

	sort.Sort(d)
	for i := 1; i < len(d.data); i++ {
		if d.data[i][0] < d.data[i-1][1] {
			return false
		}

	}

	//for i := 1; i < len(d); i++ {
	//	if d[i][0] < d[i-1][1] {
	//		return false
	//	}
	//}

	return true
}


//type conf struct {
//	data [][]int
//}
//
//
//func (d conf) Len() int {
//	return len(d.data)
//}
//
//func (d conf) Less(i, j int) bool {
//	if len(d.data[i]) == 0 {
//		return true
//	}
//	if len(d.data[j]) == 0 {
//		return false
//	}
//	return d.data[i][0] < d.data[j][0]
//}
//
//func (d conf) Swap(i, j int) {
//	d.data[i], d.data[j] = d.data[j], d.data[i]
//}

type  conf struct {
	data [][]int
}

func (d conf) Len() int {
	return len(d.data)
}

func (d conf) Less(i, j int) bool {
	if len(d.data[i]) == 0 {
		return false
	}

	if len(d.data[j]) == 0 {
		return false
	}

	return d.data[i][0] < d.data[j][0]
}

func (d conf) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}