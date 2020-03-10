package _897_increasing_order_search_tree

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

var head, newRoot *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	newRoot = &TreeNode{}
	head = newRoot
	helper(root)
	return head.Right
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}

	helper(root.Left)
	root.Left = nil	// 把left设为空
	newRoot.Right = root
	newRoot = root
	helper(root.Right)
}
