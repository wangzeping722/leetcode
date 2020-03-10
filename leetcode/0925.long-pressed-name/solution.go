package _925_long_pressed_name

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0

	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			if i == 0 {
				return false
			}
			if typed[j] == name[i-1] {
				j++
			} else {
				return false
			}
		}
	}

	if i < len(name) {
		return false
	}
	return true
}
