package sorts

/*
类似冒泡排序，通过两层嵌套的循环来遍历、比较数组中元素。
不同的是选择排序在内层循环中不交换，而是选出最小值的位置，然后在外层循环中交换最小值到头部。
最终实现数组依次按照最小值排序。
选择排序即选择最值位置排序，效率差

不稳定算法

*/

//func Selection(arr []int) {
//	for i := 0; i < len(arr)-1; i++ {
//		min := i
//		for j := i + 1; j < len(arr); j++ {
//			if arr[j] < arr[min] {
//				min = j
//			}
//		}
//		arr[min], arr[i] = arr[i], arr[min]
//	}
//}
//
//func Selection1(arr []int) {
//	for i := 0; i < len(arr)-1; i++ {
//		min := i
//		for j := i + 1; j < len(arr); j++ {
//			if arr[j] < arr[min] {
//				min = j
//			}
//		}
//		arr[min], arr[i] = arr[i], arr[min]
//	}
//}

func Selection(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func Select(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i:=0; i < len(arr)-1; i++ {
		min := i
		// 找到最小的值
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
