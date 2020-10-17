package _977

import "sort"

func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}

	sort.Ints(A)
	return A
}

// 双指针
func sortedSquares1(A []int) []int {
	n := len(A)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n-1; pos >= 0; pos-- {
		if v, w := A[i]*A[i], A[j]*A[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return A
}