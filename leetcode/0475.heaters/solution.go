package _475_heaters

import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	res := 0
	for _, house := range houses {
		left, right := 0, len(heaters)-1
		houseRes := 0
		for left < right {
			middle := left + (right-left)/2
			if heaters[middle] < house {
				left = middle + 1
			} else {
				right = middle
			}

			if heaters[left] == house {
				houseRes = 0
			} else if heaters[left] < house {
				houseRes = house - heaters[left]
			} else { // 供暖器在右侧
				if left == 0 {
					houseRes = heaters[left] - house
				} else {
					houseRes = getMin(heaters[left] - house, house - heaters[left - 1])
				}
			}

			res = getMax(res, houseRes)
		}
	}

	return res
}

func getMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
