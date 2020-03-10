package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxPath int

func maxPathSum(root *TreeNode) int {
	maxPath = math.MinInt64
	helper(root)
	return maxPath
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(0, helper(root.Left))   // 左边的最大值
	right := max(0, helper(root.Right)) // 右边的最大值

	maxPath = max(maxPath, left+right+root.Val)
	return max(left, right) + root.Val
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
