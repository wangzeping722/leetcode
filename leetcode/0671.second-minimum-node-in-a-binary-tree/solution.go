package _671_second_minimum_node_in_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 太绕了
// 使用全局变脸更加简单 TODO
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}

	left := root.Left.Val
	right := root.Right.Val

	if left == root.Val {
		left = findSecondMinimumValue(root.Left)
	}

	if right == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left != -1 && right != -1 {
		if left < right {
			return left
		}
		return right
	}
	if left == -1 {
		return right
	}
	return left
}
