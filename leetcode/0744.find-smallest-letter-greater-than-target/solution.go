package _744_find_smallest_letter_greater_than_target

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := (left + right) >> 1
		if letters[mid] > target {
			right = mid
		} else if letters[mid] <= target {
			left = mid + 1
		}
	}
	return letters[left%len(letters)]
}
