package main

func partition(head *ListNode, x int) *ListNode {
	var ltHead *ListNode
	var ltTail *ListNode
	var gtHead *ListNode
	var gtTail *ListNode
	for head != nil {
		if head.Val < x {
			if ltHead == nil {
				ltHead = head
				ltTail = head
			} else {
				ltTail.Next = head
				ltTail = head
			}
		} else {
			if gtHead == nil {
				gtHead = head
				gtTail = head
			} else {
				gtTail.Next = head
				gtTail = head
			}
		}
		head = head.Next
	}
	if gtTail != nil {
		gtTail.Next = nil
	}
	if ltHead != nil {
		ltTail.Next = gtHead
		return ltHead
	} else {
		return gtHead
	}

}
