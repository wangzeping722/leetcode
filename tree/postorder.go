package tree

//
//func postorderTraversal(root *TreeNode) []int {
//	res := []int{}
//	var dfs func(*TreeNode)
//	dfs = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//
//		dfs(root.Left)
//		dfs(root.Right)
//		res = append(res, root.Val)
//	}
//	dfs(root)
//	return res
//}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right!= nil {
			stack = append(stack, node.Right)
		}
		res = append(res, node.Val)
	}
	reverse(res)
	return res
}

func post(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		res = append(res, node.Val)
	}
	reverse(res)
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i,j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
