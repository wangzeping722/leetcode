package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	l := len(postorder)
	node := &TreeNode{Val: postorder[l-1]} //中间节点
	idx := index(inorder, postorder[l-1])
	node.Left = buildTree(inorder[:idx], postorder[:idx])
	node.Right = buildTree(inorder[idx+1:], postorder[idx:l-1])
	return node
}

func index(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	buildTree([]int{9,3,15,20,7}, []int{9,15,7,20,3})
}
