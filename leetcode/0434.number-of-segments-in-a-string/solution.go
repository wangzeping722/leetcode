package _434_number_of_segments_in_a_string

func countSegments(s string) int {
	count := 0

	for i, r := range s {
		if (i == 0 || s[i-1] == ' ') && r != ' ' {
			count++
		}
	}
	return count
}