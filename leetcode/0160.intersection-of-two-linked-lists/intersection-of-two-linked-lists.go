package problem0160

type ListNode struct {
	Val  int
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

// 具体算法如下:
// 	模拟走路:两人都必须走 len(a) + len(b)的长度,假设有相交点,那么他们最后会走相同的一段路, 我们可以得到相交的起点

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	a, b := headA, headB
//	hasLinkedToA, hasLinkedToB := false, false
//	if a != nil && b != nil {
//		if a == b {
//			return b
//		}
//
//		a, b = a.Next, b.Next
//		if a == nil && !hasLinkedToB {
//			a = headB
//			hasLinkedToB = true
//		}
//
//		if b == nil && !hasLinkedToA {
//			b = headA
//			hasLinkedToA = true
//		}
//	}
//	return nil
//}
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	a, b := headA, headB
//	linkA, linkB := false, false
//	for a != nil && b != nil {
//		if a == b {
//			return a
//		}
//
//		a, b = a.Next, b.Next
//		if a == nil && !linkB {
//			a = headB
//			linkB = true
//		}
//		if b == nil && !linkA {
//			b = headA
//			linkA = true
//		}
//	}
//	return nil
//}
