package sorts

/*
典型的分治思想：
	如果要排序一个数组，我们先把数组从中间分成前后两个部分，然后对前后两部分分别排序，
	再将排好序的两部分合并在一起
*/
func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2
	MergeSort(arr[:mid]) // 这里
	MergeSort(arr[mid:])
	merge(arr)
}

func merge(arr []int) {
	length := len(arr)
	tempArr := make([]int, length)
	left, mid := 0, length/2-1 // 这里
	right := mid + 1
	i := 0
	for left <= mid && right <= length-1 {
		if arr[left] <= arr[right] {
			tempArr[i] = arr[left]
			left++
		} else {
			tempArr[i] = arr[right]
			right++
		}
		i++
	}
	if left <= mid {
		copy(tempArr[i:], arr[left:mid+1])
	}
	if right <= length-1 {
		copy(tempArr[i:], arr[right:])
	}
	copy(arr, tempArr)
}