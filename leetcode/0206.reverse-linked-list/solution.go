package problem0206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for head != nil {
		head = head.Next
		cur.Next, prev, cur = prev, cur, head
	}
	return prev
}

