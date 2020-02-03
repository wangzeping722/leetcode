package _455_assign_cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	for i,j:=len(g)-1,len(s)-1; i >= 0; i-- {
		if j == -1 {
			return res
		}
		if g[i] <= s[j] {
			j--
			res++
		}
	}
	return res
}