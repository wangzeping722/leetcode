package _382

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	res := 0
	i := 0
	for this.head != nil {
		i++
		if rand.Intn(i) == 0 {
			res = this.head.Val
		}
		this.head = this.head.Next
	}
	return res
}
