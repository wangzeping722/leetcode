package problem0112

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Problem0112(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(true, hasPathSum(&TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}, 0))
}
