package problem0111

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 没有字节点的node才是叶子节点
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



