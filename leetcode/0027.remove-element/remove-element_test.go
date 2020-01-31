package problem0027

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

func Test_removeElement(t *testing.T) {
	ast := assert.New(t)

	// 元素顺序可以改变
	qs := []question{
		{
			para{
				one: []int{1, 2, 3, 3, 4, 3},
				two: 3,
			}, ans{one: []int{1, 2, 4}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, p.one[:removeElement(p.one, p.two)], "输入:%v", p)
	}

}
