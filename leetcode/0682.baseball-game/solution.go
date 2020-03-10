package _682_baseball_game

import "strconv"

func calPoints(ops []string) int {
	arr := make([]int, len(ops))
	i := 0
	for _, s := range ops {
		switch s {
		case "+":
			arr[i] = arr[i-1] + arr[i-2]
			i++
		case "D":
			arr[i] = 2*arr[i-1]
			i++
		case "C":
			arr[i-1] = 0
			i--
		default:
			arr[i], _ = strconv.Atoi(s)
			i++
		}
	}
	sum := 0
	for j:=0; j < len(arr); j++ {
		sum += arr[j]
	}
	return sum
}
