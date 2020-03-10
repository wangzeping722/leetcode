package _543_diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	helper(root, 0)
	return max
}

func helper(root *TreeNode, height int) int {
	if root == nil {
		return 0
	}
	leftHeight := helper(root.Left, height)
	rightHeight := helper(root.Right, height)

	length := leftHeight + rightHeight
	if length > max {
		max = length
	}
	return maxNum(leftHeight, rightHeight) + 1
}

func maxNum(i, j int) int {
	if i > j {
		return i
	}
	return j
}
