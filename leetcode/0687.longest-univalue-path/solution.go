package _687_longest_univalue_path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	helper(root)
	return res
}


func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := helper(root.Left), helper(root.Right)
	if root.Left != nil && root.Val == root.Left.Val {
		left += 1
	} else {
		left = 0
	}

	if root.Right != nil && root.Val == root.Right.Val {
		right += 1
	} else {
		right = 0
	}
	res = max(res, left+right)
	return max(left, right)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}