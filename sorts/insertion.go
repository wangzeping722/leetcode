package sorts

/*
插入排序将数组分为两个区域，一个是排序区域，一个是非排序区域，
初始已排序区域只有一个元素，就是数组的第一个元素。
插入算法的核心思想：
	取未排序区间的元素，在已排序区间中找到合适的位置插入进去，并保证已排序区域一直有序，
	重复这个过程，直到未排序区间中的元素为空，算法结束。

原地排序算法

稳定算法
*/

//func Insertion(arr []int) []int {
//	for i := 1; i < len(arr); i++ {
//		val := arr[i]
//		j := i - 1
//		for ; j >= 0 && arr[j] > val; j-- {
//			arr[j+1] = arr[j]
//		}
//		arr[j+1] = val
//	}
//	return arr
//}
//
//func Insertion1(arr []int) {
//	for i := 1; i < len(arr); i++ {
//		temp := arr[i]
//		j := i - 1
//		for ; j >= 0 && arr[j] > temp; j-- {
//			arr[j+1] = arr[j]
//		}
//		arr[j+1] = temp
//	}
//}

// 稳定排序
func Insertion(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		tempVal := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > tempVal; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = tempVal
	}
}

func Insert(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		tempVal := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > tempVal; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = tempVal
	}
}
