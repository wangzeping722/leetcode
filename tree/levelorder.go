package tree

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curLen := len(queue)
		curArr := make([]int, curLen)
		for i := 0; i < curLen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			curArr[i] = queue[i].Val
		}
		queue = queue[curLen:]
		res = append(res, curArr)
	}
	return res
}
