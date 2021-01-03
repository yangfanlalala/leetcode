package main

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	p := 0
//	var head *ListNode
//	var current *ListNode
//	for l1!=nil || l2!=nil {
//		i, j, rst := 0, 0, 0
//		if l1 != nil {
//			i = l1.Val
//		}
//		if l2 != nil {
//			j = l2.Val
//		}
//		rst = i + j + p
//		if rst >= 10 {
//			rst -= 10
//			p = 1
//		} else {
//			p = 0
//		}
//		if head == nil {
//			head = &ListNode{
//				Val:  rst,
//				Next: nil,
//			}
//			current = head
//		} else {
//			current.Next = &ListNode{
//				Val:  rst,
//				Next: nil,
//			}
//			current = current.Next
//		}
//		if l1 != nil {
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			l2 = l2.Next
//		}
//	}
//	if p > 0 {
//		current.Next = &ListNode{
//			Val: p,
//		}
//	}
//	return head
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1, len2 := 0, 0
	l1h, l2h, p := l1, l2, 0
	var tail *ListNode
	for l1h != nil {
		len1++
		l1h = l1h.Next
	}
	for l2h != nil {
		len2++
		l2h = l2h.Next
	}
	if len2 > len1 {
		l1, l2 = l2, l1
	}
	l1h, l2h = l1, l2
	for l1h != nil {
		rst, i, j := 0, 0, 0
		i = l1h.Val
		if l2h != nil {
			j = l2h.Val
		}
		rst = i + j + p
		if rst >= 10 {
			rst -= 10
			p = 1
		} else {
			p = 0
		}
		l1h.Val = rst
		if l2h != nil {
			l2h = l2h.Next
		}
		if l1h.Next == nil {
			tail = l1h
		}
		l1h = l1h.Next
	}
	if p > 0 {
		tail.Next = &ListNode{
			Val:  p,
			Next: nil,
		}
	}
	return l1
}