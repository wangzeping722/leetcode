package _257_binary_tree_paths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return dfs("", []string{}, root)
}

func dfs(path string, paths []string, root *TreeNode) []string {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		paths = append(paths, path)
	} else {
		path += "->"
		if root.Left != nil {
			paths = dfs(path, paths, root.Left)
		}
		if root.Right != nil {
			paths = dfs(path, paths, root.Right)
		}
	}

	return paths
}
