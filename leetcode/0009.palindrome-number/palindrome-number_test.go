package problem0009

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{one: 12321},
			a: ans{true},
		},
		{
			p: para{-11},
			a: ans{false},
		},
	}

	for _, q := range qs {
		p, a := q.p, q.a
		ast.Equal(a.one, isPalindrome_str(p.one), "输入:%v", p)
	}
}
