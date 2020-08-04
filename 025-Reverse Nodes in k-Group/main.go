package main

import "fmt"
type ListNode struct {
	Val 	int
	Next	*ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newNode := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newNode
}

func reverse(first, end *ListNode) *ListNode {
	var prev *ListNode = end
	
	for first != end{
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}


func main() {
	newListNode := func(values []int) *ListNode {
		var l, next *ListNode
		for _, v := range values {
			if l == nil {
				l = &ListNode{Val: v}
				next = l
				continue
			}
			next.Next = &ListNode{Val: v}
			next = next.Next
		}

		return l
	}


	l1 := newListNode([]int{1, 2, 3, 4, 5})
	l1 = reverseKGroup(l1, 2)
	for l1 != nil{
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}