package _717_1_bit_and_2_bit_characters

func isOneBitCharacter(bits []int) bool {
	i := len(bits) - 2
	for i >= 0 && bits[i] > 0 {
		i--
	}
	return (len(bits)-i)%2 == 0
}

func main() {
	isOneBitCharacter([]int{1, 0, 0})
}
