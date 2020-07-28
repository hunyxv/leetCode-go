package main

import "fmt"


type ListNode struct {
	Val int
   	Next *ListNode
}

func margeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val{
		l1.Next = margeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = margeTwoLists(l1, l2.Next)
	return l2
}


func main() {
	var l1, l2, next *ListNode
	for _, v := range([]int{1, 2, 4}) {
		if l1 == nil {
			l1 = &ListNode{Val: v}
			next = l1
			continue
		}
		next.Next = &ListNode{Val: v}
		next = next.Next
	}

	for _, v := range([]int{1, 3, 4}) {
		if l2 == nil {
			l2 = &ListNode{Val: v}
			next = l2
			continue
		}
		next.Next = &ListNode{Val: v}
		next = next.Next
	}

	newList := margeTwoLists(l1, l2)
	for newList != nil{
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}