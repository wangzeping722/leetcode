package _538_convert_bst_to_greater_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	convert(root, 0)
	return root
}

func convert(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	sum = convert(root.Right, sum)
	root.Val += sum
	return convert(root.Left, root.Val)
}
