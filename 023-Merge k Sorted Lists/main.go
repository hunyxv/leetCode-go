package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func mergeKLists(list []*ListNode) *ListNode {
	if len(list) == 0{
		return nil
	}

	if len(list) == 1 {
		return list[0]
	}

	n := len(list) / 2
	left := mergeKLists(list[:n])
	right := mergeKLists(list[n:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Val < right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	}
	right.Next = mergeTwoLists(left, right.Next)
	return right
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


	l1 := newListNode([]int{1, 2, 4})
	l2 := newListNode([]int{1, 3, 4})
	l3 := newListNode([]int{2, 5, 6})

	newList := mergeKLists([]*ListNode{l1, l2, l3})
	for newList != nil{
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}