package _783_minimum_distance_between_bst_nodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var minInt int
var preNode *TreeNode

func minDiffInBST(root *TreeNode) int {
	minInt = math.MaxInt64
	preNode = nil
	helper(root)
	return minInt
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	if preNode != nil {
		minInt = min(minInt, abs(root.Val, preNode.Val))
	}
	preNode = root
	helper(root.Right)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
