package _257

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []string
var node []string

func binaryTreePaths(root *TreeNode) []string {
	res = make([]string, 0)
	node = make([]string, 0)
	if root == nil {
		return res
	}

	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		node = append(node, strconv.Itoa(root.Val))
		if root.Left == nil && root.Right == nil {
			res = append(res, strings.Join(node, "->"))
		}else {
			dfs(root.Left)
			dfs(root.Right)
		}
		node = node[:len(node)-1]
	}
}
