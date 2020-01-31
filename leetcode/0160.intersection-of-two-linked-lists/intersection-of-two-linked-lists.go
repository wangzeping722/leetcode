package problem0160

type ListNode struct {
	Val int
	Next *ListNode
}

// TODO 没想出解法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	hasLinkedToA, hasLinkedToB := false, false
	for a != nil && b != nil {
		if a == b {
			return b
		}

		a, b = a.Next, b.Next
		if a == nil && !hasLinkedToB {
			a = headB
			hasLinkedToB = true
		}
		if b == nil && !hasLinkedToA {
			b = headA
			hasLinkedToA = true
		}
	}

	return nil
}
