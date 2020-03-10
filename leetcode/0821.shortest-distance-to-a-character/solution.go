package _821_shortest_distance_to_a_character

import "math"

func shortestToChar(S string, C byte) []int {
	res := make([]int, len(S))
	for i := 0; i < len(res); i++ {
		res[i] = math.MaxInt64
	}
	left := -1
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			res[i] = 0
			left = i
		} else if left >= 0 {
			res[i] = min(res[i], i-left)
		}
	}
	right := -1
	for i := len(S)-1; i >= 0; i-- {
		if S[i] == C {
			right = i
		} else if right >= 0 {
			res[i] = min(res[i], right-i)
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
