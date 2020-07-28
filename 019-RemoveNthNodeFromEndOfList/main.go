package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p1, p2 , tmp *ListNode = head, head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for {
		if p2 == nil {
			if tmp == head {
				return p1.Next
			}
			tmp.Next = p1.Next
			return head
		}
		tmp = p1
		p1 = p1.Next
		p2 = p2.Next
	}
}

func main() {
	head := &ListNode{}
	p := head
	for i := 1; i < 20; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}

	h2 := removeNthFromEnd(head, 20)
	for h2 != nil {
		fmt.Println(h2.Val)
		h2 = h2.Next
	}
}