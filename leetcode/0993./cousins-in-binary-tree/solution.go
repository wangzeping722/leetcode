package cousins_in_binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var (
	px int
	py int
	hx int
	hy int
)
func isCousins(root *TreeNode, x int, y int) bool {
	px, py, hx, hy = 0, 0, 0, 0
	dfs(root, x, y, 0)
	return hx == hy &&  px != py
}


func dfs(root *TreeNode, x, y, depth int) {
	if root == nil {
		return
	}
	if px != 0 && py != 0 {
		return
	}

	if root.Left != nil {
		if root.Left.Val == x {
			px = root.Val
			hx = depth+1
		} else if root.Left.Val == y {
			py = root.Val
			hy = depth+1
		}
	}


	if root.Right != nil {
		if root.Right.Val == x {
			px = root.Val
			hx = depth+1
		} else if root.Right.Val == y {
			py = root.Val
			hy = depth+1
		}
	}

	dfs(root.Left, x, y ,depth+1)
	dfs(root.Right, x, y ,depth+1)
}