package sorts

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array1 := random()
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)
	quickSort(array1)
	array2.Sort()
	fmt.Println(array1)
	fmt.Println(array2)
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}