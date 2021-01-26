package addtwonumbers

import "testing"

func Test(t *testing.T) {
	_l1 := &ListNode{}
	_l2 := &ListNode{}
	l1, l2 := _l1, _l2
	for _, v := range []int{2, 4, 3, 5} {
		_l1.Next = &ListNode{Val: v}
		_l1 = _l1.Next
	}
	for _, v := range []int{5, 6, 4, 5, 9} {
		_l2.Next = &ListNode{Val: v}
		_l2 = _l2.Next
	}
	res := addTwoNumbers(l1.Next, l2.Next)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
