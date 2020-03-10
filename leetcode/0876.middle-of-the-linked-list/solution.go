package _876_middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0
	h := head
	for h != nil {
		count++
		h = h.Next
	}
	pos := count / 2
	for pos > 0 {
		head = head.Next
		pos--
	}
	return head
}


func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}