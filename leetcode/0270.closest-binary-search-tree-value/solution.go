package _270_closest_binary_search_tree_value

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	var min float64 = math.MaxFloat64
	minVal := 0
	for root != nil {
		temp := math.Abs(float64(root.Val) - target)
		if temp < min {
			min = temp
			minVal = root.Val
		}

		if float64(root.Val) > target {
			root = root.Left
		} else if float64(root.Val) < target {
			root = root.Right
		}  else {
			return root.Val
		}
	}
	return minVal
}
