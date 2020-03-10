package _617_merge_two_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	s := []*TreeNode{}
	s = append(s, t1, t2)
	for len(s) != 0 {
		r1 := s[0]
		r2 := s[1]
		s = s[2:]
		r1.Val += r2.Val
		if r1.Left != nil && r2.Left != nil {
			s = append(s, r1.Left, r2.Left)
		} else if r1.Left == nil {
			r1.Left = r2.Left
		}
		if r1.Right!= nil && r2.Right!= nil {
			s = append(s, r1.Right, r2.Right)
		} else if r1.Right == nil {
			r1.Right = r2.Right
		}
	}
	return t1
}
