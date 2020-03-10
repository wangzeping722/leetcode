package _849_maximize_distance_to_closest_person

import "math"

// 双指针遍历
func maxDistToClosest(seats []int) int {
	dists := make([]int, len(seats))
	for i := 0; i < len(dists); i++ {
		dists[i] = math.MaxInt64
	}
	for i, left := 0, -1; i < len(seats); i++ {
		if seats[i] == 1 {
			left = i
			dists[i] = 0
		} else if left >= 0 && seats[i] == 0 {
			dists[i] = i - left
		}
	}
	for i, right := len(seats)-1, -1; i >= 0; i-- {
		if seats[i] == 1 {
			right = i
		} else if right >= 0 && seats[i] == 0 {
			dists[i] = min(dists[i], right-i)
		}
	}
	maxDist := 0
	for i:=0;i<len(dists);i++ {
		if dists[i] > maxDist {
			maxDist = dists[i]
		}
	}
	return maxDist
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}


func maxDistToClosest1(seats []int) int {
	l := len(seats)
	k := -1
	res := 0
	for i := 0; i < l; i++ {
		if seats[i] == 1 {
			if k >= 0 {
				res = max(res, (i-k)/2)
			} else {
				res = i
			}
			k = i
		}
	}
	res = max(res, len(seats)-1-k)
	return res
}

func max(i, j int ) int {
	if i > j {
		return i
	}
	return j
}