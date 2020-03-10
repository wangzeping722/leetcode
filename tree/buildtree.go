package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return node
	}

	idx := index(inorder, preorder[0])
	node.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
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
