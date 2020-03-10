package sorts

/*
典型的分治思想：
	如果要排序一个数组，我们先把数组从中间分成前后两个部分，然后对前后两部分分别排序，
	再将排好序的两部分合并在一起
*/

//func MergeSort(arr []int) {
//	if len(arr) <= 1 {
//		return
//	}
//	mergeSort(arr, 0, len(arr)-1)
//}
//
//func mergeSort(arr []int, low, high int) {
//	if low >= high {
//		return
//	}
//
//	mid := low + (high-low)>>1
//	mergeSort(arr, low, mid)
//	mergeSort(arr, mid+1, high)
//	merge(arr, low, mid, high)
//}
//
//func merge(arr []int, low, mid, high int) {
//	length := high - low + 1
//	tempArr := make([]int, length)
//
//	i, ileft, iright := 0, low, mid+1
//	for ileft <= mid && iright <= high {
//		if arr[ileft] <= arr[iright] {
//			tempArr[i] = arr[ileft]
//			ileft++
//		} else {
//			tempArr[i] = arr[iright]
//			iright++
//		}
//		i++
//	}
//
//	// copy the rest data
//	if ileft <= mid {
//		copy(tempArr[i:], arr[ileft:mid+1])
//	}
//	if iright <= high {
//		copy(tempArr[i:], arr[iright:high+1])
//	}
//	copy(arr[low:high+1], tempArr[:])
//}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	merge(arr)
}

func merge(arr []int) {
	length := len(arr)
	tempArr := make([]int, length)
	left, mid := 0, length/2-1
	right := mid+1
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
		copy(tempArr[i:], arr[right:length])
	}
	copy(arr, tempArr)
}
