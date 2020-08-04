package main

import "fmt"

type ListNode struct {
	Val 	int
	Next	*ListNode
}

func swapNode(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}

	if list.Next == nil {
		return list
	}

	tmp := list
	list = list.Next
	tmp.Next = swapNode(list.Next)
	list.Next = tmp

	return list
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


	l1 := newListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	l1 = swapNode(l1)
	for l1 != nil{
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
}