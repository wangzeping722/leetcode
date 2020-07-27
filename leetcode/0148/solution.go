package _148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	tmp := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(tmp)
	h := new(ListNode)
	res := h
	for left != nil && right != nil {
		if left.Val <= right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}

	return res.Next
}
