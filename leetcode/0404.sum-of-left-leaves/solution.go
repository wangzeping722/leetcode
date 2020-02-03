package _404_sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil{
		left = root.Left.Val
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right) + left
}
