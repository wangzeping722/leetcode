package _832_flipping_an_image

func flipAndInvertImage(A [][]int) [][]int {
	for _, arr := range A {
		for i, j := 0, len(arr)-1; i <j; i,j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		for i:=0; i < len(arr); i++ {
			arr[i] ^= 1
		}
	}
	return A
}
