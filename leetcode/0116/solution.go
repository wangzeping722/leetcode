package _116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTowNode(root.Left, root.Right)
	return root
}

func connectTowNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2
	connectTowNode(node1.Left, node1.Right)
	connectTowNode(node2.Left, node2.Right)
	connectTowNode(node1.Right, node2.Left)
}