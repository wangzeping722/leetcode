package _914_x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, num := range deck {
		m[num]++
	}
	g := -1
	for _, v := range m {
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

// 辗转相除法
func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

func gcd1(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd1(y, x%y)
}