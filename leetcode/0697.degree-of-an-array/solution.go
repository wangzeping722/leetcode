package _696_count_binary_substrings

func findShortestSubArray(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	s := []int{}
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			s = []int{k}
		} else if v == max {
			s = append(s, k)
		}
	}

	minLen := len(nums)
	for _, k := range s {
		count, left := 0, 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == k {
				if count == 0 {
					left = i
				}
				count++
				if count == max {
					if i-left+1 < minLen {
						minLen = i-left+1
					}
				}
			}
		}
	}
	return minLen
}
