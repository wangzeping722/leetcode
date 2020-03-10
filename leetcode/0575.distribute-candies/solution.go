package _575_distribute_candies

func distributeCandies(candies []int) int {
	m := make(map[int]bool, len(candies))
	for _, candy := range candies {
		m[candy] = true
	}
	if len(m) < len(candies)/2 {
		return len(m)
	}
	return len(candies)/2
}
