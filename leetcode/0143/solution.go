package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	stack := make([]*ListNode, 0)
	slow = slow.Next
	for slow != nil {
		stack = append(stack, slow)
		slow = slow.Next
	}

	cur := head
	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		node.Next = cur.Next
		cur.Next = node
		cur = node.Next
	}
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
