package _443_string_compression

import "strconv"

func compress(chars []byte) int {
	left, size := 0, 0

	for right := 0; right <= len(chars); right++ {
		if right == len(chars) || chars[right] != chars[left] {
			chars[size] = chars[left]
			size++
			if right-left > 1 {
				for _, r := range strconv.Itoa(right - left) {
					chars[size] = byte(r)
					size++
				}
			}
			left = right
		}
	}

	return size
}
