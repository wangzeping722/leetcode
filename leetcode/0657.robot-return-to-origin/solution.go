package _657_robot_return_to_origin

func judgeCircle(moves string) bool {
	u, d, l, r := 0,0,0,0

	for _, b := range moves {
		switch b {
		case 'U':
			u++
		case 'D':
			d++
		case 'L':
			l++
		case 'R':
			r++
		}
	}

	if u != d || l != r {
		return false
	}
	return true
}
