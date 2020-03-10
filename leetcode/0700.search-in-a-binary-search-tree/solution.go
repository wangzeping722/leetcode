package _700_search_in_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if  root == nil || root.Val == val {
		return root
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
