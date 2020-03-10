package _532_k_diff_pairs_in_an_array

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	numsHash := make(map[int]bool)
	diffHash := make(map[int]bool)

	for _, num := range nums {
		if numsHash[num-k] {
			diffHash[num-k] = true
		}
		if numsHash[num+k] {
			diffHash[num] = true
		}
		numsHash[num] = true
	}

	return len(diffHash)
}