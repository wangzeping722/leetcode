package _653_two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	return helper(root, m, k)
}

func helper(root *TreeNode, m map[int]bool, k int) bool {
	if root == nil {
		return false
	}
	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true
	return helper(root.Left, m, k) || helper(root.Right, m, k)
}
