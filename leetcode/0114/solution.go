package _114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	// 拉平左右子树
	flatten(root.Left)
	flatten(root.Right)
	// 左右子树已经被拉平
	left := root.Left
	right := root.Right

	// 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 将右子树接到当前左子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}


