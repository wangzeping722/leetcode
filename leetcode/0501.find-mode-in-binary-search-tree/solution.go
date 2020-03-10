package _501_find_mode_in_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 多个测试用例要使用闭包
var (
	res          []int
	currentCount int
	maxCount     int
	preNode      *TreeNode
)

func findMode(root *TreeNode) []int {
	res = []int{}
	currentCount = 0
	maxCount = 0
	preNode = nil
	inOrder(root)
	return res
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	if preNode != nil && root.Val == preNode.Val {
		currentCount++
	} else {
		currentCount = 1
	}

	if currentCount == maxCount {
		res = append(res, root.Val)
	} else if currentCount > maxCount {
		maxCount = currentCount
		res = []int{root.Val}
	}
	preNode=root
	inOrder(root.Right)
}
