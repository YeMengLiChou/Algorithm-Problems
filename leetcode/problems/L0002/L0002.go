package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := (*ListNode)(nil)
	res := &root
	cnt := 0
	for l1 != nil || l2 != nil {
		sum := cnt
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cnt = sum / 10
		sum %= 10
		(*res) = &ListNode{Val: sum}
		res = &(*res).Next
	}
	if cnt > 0 {
		(*res) = &ListNode{Val: cnt}
		res = &(*res).Next
	}
	return root
}
