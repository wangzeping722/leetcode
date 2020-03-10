package _872_leaf_similar_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1 := []int{}
	s2 := []int{}
	helper(root1, &s1)
	helper(root2, &s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func helper(root *TreeNode, s *[]int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*s = append(*s, root.Val)
	}
	helper(root.Left, s)
	helper(root.Right, s)
}
