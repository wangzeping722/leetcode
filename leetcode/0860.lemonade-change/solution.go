package _860_lemonade_change

func lemonadeChange(bills []int) bool {
	d5, d10 := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			d5++
		case 10:
			d10++
			d5--
		case 20:
			if d10 >= 1 {
				d10--
				d5--
			} else {
				d5 -= 3
			}
		}
		if d5 < 0 {
			return false
		}
	}
	return true
}