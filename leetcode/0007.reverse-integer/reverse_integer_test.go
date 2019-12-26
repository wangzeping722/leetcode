package problem0007

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		{
			p: para{one:15},
			a: ans{51},
		},
		{
			p: para{one:-150},
			a: ans{-51},
		},
	}


	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverse(p.one), "输入:%v", p)
	}
}