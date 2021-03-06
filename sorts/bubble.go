package sorts

/*
冒泡排序：
	冒泡排序只会操作相邻的两个数据，每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系，
	如果不满足就互换位置。
*/
func Bubble(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := len(arr) - 1; i > 0; i-- {
		flag := true
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j-1] {
				flag = false
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		if flag {
			return
		}
	}
}

func bubble1(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := len(arr)-1; i > 0; i-- {
		flag := true
		for j := 1; j <= i; j++ {
			if arr[j] < arr[j-1] {
				flag = false
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		if flag {
			return
		}
	}
}