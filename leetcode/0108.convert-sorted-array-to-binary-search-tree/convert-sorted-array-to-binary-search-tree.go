package problem0108

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//TODO
/*
1. 选取中间节点
2. 构造该节点的左子树
3. 构造该节点的右子树
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums)/2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
