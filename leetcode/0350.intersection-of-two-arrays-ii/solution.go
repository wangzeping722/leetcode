package _350_intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	mn1 := make(map[int]int, len(nums1))
	mn2 := make(map[int]int, len(nums2))
	res := []int{}
	for _, num := range nums1 {
		mn1[num]++
	}

	for _, num := range nums2 {
		mn2[num]++
	}

	for key, val1 := range mn1 {
		if val2 := mn2[key]; val2 > 0{
			v := val1
			if val1 > val2 {
				v = val2
			}
			for i :=0 ;i < v; i++ {
				res = append(res, key)
			}
		}
	}
	return res
}