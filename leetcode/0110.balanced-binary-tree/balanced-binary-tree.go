package problem0110

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := recur(root)
	return balanced
}

func recur(root *TreeNode)(int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftBalanced := recur(root.Left)
	rightDepth, rightBalanced := recur(root.Right)
	x := leftDepth - rightDepth
	if leftBalanced && rightBalanced && x <= 1 && x >= -1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}