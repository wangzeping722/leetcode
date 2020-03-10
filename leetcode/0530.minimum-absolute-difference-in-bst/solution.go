package _530_minimum_absolute_difference_in_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preNode *TreeNode

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return math.MaxInt64
	}

	leftMin := getMinimumDifference(root.Left)
	if preNode != nil {
		temp := int(math.Abs(float64(preNode.Val - root.Val)))
		if temp < leftMin{
			leftMin = temp
		}
	}
	preNode = root
	rightMin := getMinimumDifference(root.Right)
	if leftMin < rightMin {
		return leftMin
	}
	return rightMin
}
