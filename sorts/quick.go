package sorts

/*
快排也是利用了分治的思想
	如果要排序数组中下标从p到r之间的一组数据, 我们选择p到r之间的任意一个数据
作为pivot（分区点）
我们遍历p到r之间的数据，将小于pivot的数据放到左边，将大于pivot的数据放到右边，
等于pivot的数据放到中间。经过这一步骤，数组p到r之间的数据就被分成了三部分，
前面p到q-1之间都是小于pivot的，中间是等于pivot，后面q+1到r是大于pivot
*/
//func QuickSort(arr []int) {
//	if len(arr) <= 1 {
//		return
//	}
//
//	p := arr[0]
//	head, tail := 0, len(arr)-1
//	for i := 1; i <= tail; {
//		if arr[i] > p {
//			arr[i], arr[tail] = arr[tail], arr[i]
//			tail--
//		} else if arr[i] < p {
//			arr[i], arr[head] = arr[head], arr[i]
//			head++
//			i++
//		} else {
//			i++
//		}
//		//fmt.Printf("head: %d, i: %d\n", head, i)
//	}
//	QuickSort(arr[:head])
//	QuickSort(arr[tail+1:])
//}

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] > p {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else if arr[i] < p {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		} else {
			i++
		}
	}

	QuickSort(arr[:left])
	QuickSort(arr[right+1:])
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := arr[0]
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[i] > p {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else if arr[i] < p {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		} else {
			i++
		}
	}

	quickSort(arr[:left])
	quickSort(arr[right+1:])
}