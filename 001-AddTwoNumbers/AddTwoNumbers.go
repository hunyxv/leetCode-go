package addtwonumbers

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	head := &ListNode{}
	rhead := head
	var tmp int
	var sum int
	for l1 != nil || l2 != nil {

		if l1 == nil {
			sum = tmp + l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			sum = tmp + l1.Val
			l1 = l1.Next
		} else {
			sum = tmp + l1.Val + l2.Val
			l1, l2 = l1.Next, l2.Next
		}

		tmp = sum / 10
		v := sum % 10
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	if tmp != 0 {
		head.Next = &ListNode{Val: tmp}
	}
	return rhead.Next
}
