package tree


func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}


func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	for len(stack) != 0 || cur != nil{
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		cur = top.Right
	}
	return res
}