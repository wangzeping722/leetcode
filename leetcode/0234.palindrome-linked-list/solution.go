package _234_palindrome_linked_list

type ListNode struct {
	Val int
	Next *ListNode
}


func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	nums := []int{}
	for fast != nil && fast.Next != nil {
		nums = append(nums, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}
	l := len(nums) - 1
	for slow != nil {
		if nums[l] != slow.Val {
			return false
		}

		l--
		slow = slow.Next
	}
	return true
}