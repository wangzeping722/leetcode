package problem0107

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var resLevelOrderBottom [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	resLevelOrderBottom = resLevelOrderBottom[:0]
	levelOrderBottomProcess(root)
	return resLevelOrderBottom
}


func levelOrderBottomProcess(root *TreeNode) {
	if root == nil {
		return
	}

	var queue []*TreeNode
	// 模拟队列
	queue = append(queue, root)
	for len(queue) != 0 {
		curLen := len(queue)
		oneRes := make([]int, 0, curLen)
		for curLen != 0 {
			curNode := queue[0]
			queue = queue[1:]
			oneRes = append(oneRes, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			curLen--
		}
		resLevelOrderBottom = append(resLevelOrderBottom, oneRes)
	}

	var start = 0
	var end = len(resLevelOrderBottom) - 1
	for ; start < end; start, end = start+1, end-1 {
		resLevelOrderBottom[start], resLevelOrderBottom[end] = resLevelOrderBottom[end], resLevelOrderBottom[start]
	}
}
