package main

import "fmt"

type Link struct {
	val 	int
	next 	*Link
}



func addTwoNumbers(l1, l2 *Link) *Link {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &Link{}
	rhead := head
	tmp := 0
	for l1 != nil || l2 != nil {
		var sum int
		if l1 == nil {
			sum = tmp + l2.val
			l2 = l2.next
		} else if l2 == nil {
			sum = tmp + l1.val
			l1 = l1.next
		} else {
			sum = tmp + l1.val + l2.val
			l1, l2 = l1.next, l2.next
		}
		
		tmp = sum / 10
		v := sum % 10
		head.next = &Link{val:v}
		head = head.next
	}
	return rhead.next
}

func main() {
	_l1 := &Link{}
	_l2 := &Link{}
	l1, l2 := _l1, _l2
	for _, v := range []int{2, 4, 3,4} {
		_l1.next = &Link{val: v}
		_l1 = _l1.next
	}
	for _, v := range []int{5, 6, 4,5,9} {
		_l2.next = &Link{val: v}
		_l2 = _l2.next
	}

	res := addTwoNumbers(l1.next, l2.next)
	for res != nil {
		fmt.Println(res.val)
		res = res.next
	}
}