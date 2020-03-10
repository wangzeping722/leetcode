package _605_can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i:=0; i<len(flowerbed); i++ {
		if flowerbed[i] == 0 &&
			(i-1 == -1 || flowerbed[i-1] == 0) &&
			(i+1 == len(flowerbed) || flowerbed[i+1] == 0) {
			n--
			flowerbed[i] = 1
		}
	}
	return n <= 0
}