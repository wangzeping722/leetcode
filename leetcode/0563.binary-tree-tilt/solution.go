package _563_binary_tree_tilt

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	helper(root, &tilt)
	return tilt
}

// helper返回节点和
// tilt保存节点的坡度
func helper(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left, tilt)
	right := helper(root.Right, tilt)
	*tilt += int(math.Abs(float64(left - right)))
	return left + right + root.Val
}
