package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归实现
var res []int

func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	preorder(root)
	return res
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}

// 迭代实现

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	buf := []*TreeNode{root}
	for len(buf) > 0 {
		p := buf[len(buf)-1]
		buf = buf[0 : len(buf)-1]
		if p.Right != nil {
			buf = append(buf, p.Right)
		}
		if p.Left != nil {
			buf = append(buf, p.Left)
		}
		ret = append(ret, p.Val)
	}
	return ret
}

func preorderTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		ret = append(ret, p.Val)
	}
	return ret
}