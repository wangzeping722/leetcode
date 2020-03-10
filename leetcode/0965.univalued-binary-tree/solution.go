package _965_univalued_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root, root.Val)
}

func helper(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}

	if root.Val != val {
		return false
	}

	return helper(root.Left, val) && helper(root.Right, val)
}
