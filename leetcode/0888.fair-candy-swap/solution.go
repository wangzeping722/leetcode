package _888_fair_candy_swap

func fairCandySwap(A []int, B []int) []int {
	countA, countB := 0, 0
	mb := make(map[int]bool)
	for _, a := range A {
		countA += a
	}
	for _, b := range B {
		countB += b
		mb[b] = true
	}

	delta := (countB - countA) / 2
	res := make([]int, 2)
	for _, a := range A {
		if mb[delta+a] {
			res[0], res[1] = a, delta+a
			return res
		}
	}
	return nil
}
