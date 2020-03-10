package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		sum := 0
		for _, node := range queue {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(l))
		queue = queue[l:]
	}
	return res
}

func main() {
	test := []int{1,23}
	for _, num := range test {
		test = append(test, 0)
		println(num)
	}
}