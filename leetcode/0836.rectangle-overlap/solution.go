package _836_rectangle_overlap

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] ||	// in the left
		rec1[3] <= rec2[1] ||		// in the bottom
		rec1[0] >= rec2[2] ||		// in the right
		rec1[1] >= rec2[3])			// in the top
}
