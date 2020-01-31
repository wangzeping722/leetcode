package problem0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	mn := m+n-1
	m, n = m-1, n-1
	for ; m >= 0 && n >= 0; mn-- {
		if nums1[m] >= nums2[n] {
			nums1[mn] = nums1[m]
			m--
		} else {
			nums1[mn] = nums2[n]
			n--
		}
	}
	if n >= 0 {
		for ; n >= 0; n-- {
			nums1[n] = nums2[n]
		}
	}
}
