package problem0021

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode = new(ListNode)
	next := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			next.Next = l1
			l1 = l1.Next
		} else {
			next.Next = l2
			l2 = l2.Next
		}
		next = next.Next
	}
	if l1 != nil {
		next.Next = l1
	} else {
		next.Next = l2
	}
	return head.Next
}


