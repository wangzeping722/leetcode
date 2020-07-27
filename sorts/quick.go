package sorts

/*
快排也是利用了分治的思想
	如果要排序数组中下标从p到r之间的一组数据, 我们选择p到r之间的任意一个数据
作为pivot（分区点）
我们遍历p到r之间的数据，将小于pivot的数据放到左边，将大于pivot的数据放到右边，
等于pivot的数据放到中间。经过这一步骤，数组p到r之间的数据就被分成了三部分，
前面p到q-1之间都是小于pivot的，中间是等于pivot，后面q+1到r是大于pivot
*/
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	p := arr[0] // 选择第一个为基准
	left, right := 0, len(arr)-1
	for i := 1; i <= right; {
		if arr[right] > p {
			right--
			continue
		}

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

func quick(arr []int) {
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

	quick(arr[:left])
	quick(arr[right+1:])
}
