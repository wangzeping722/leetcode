package _349_intersection_of_two_arrays


func intersection(nums1 []int, nums2 []int) []int {
	mn1 := make(map[int]bool, len(nums1))
	mn2 := make(map[int]bool, len(nums2))
	res := []int{}
	for _, num := range nums1 {
		mn1[num] = true
	}

	for _, num := range nums2 {
		mn2[num] = true
	}

	for key := range mn1 {
		if _, ok := mn2[key]; ok{
			res = append(res, key)
		}
	}
	return res
}
