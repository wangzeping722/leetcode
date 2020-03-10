package _572_subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t != nil {
		return false
	}
	return  sameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func sameTree(s *TreeNode, t*TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s.Val == t.Val && sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)
}
