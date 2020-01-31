package problem0026

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{1, 1, 2},
		2,
	},

	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		5,
	},
}

func Test_removeDuplicates(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, removeDuplicates(tc.nums), "输入:%v", tc)
	}
}
