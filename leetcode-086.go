package main

func partition(head *ListNode, x int) *ListNode {
	smallHead, largeHead := &ListNode{}, &ListNode{}
	var small *ListNode = smallHead
	var large *ListNode = largeHead
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	small.Next = largeHead.Next
	large.Next = nil
	return smallHead.Next
}
