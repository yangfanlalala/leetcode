package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Next: head}
	f, s := head, newHead
	for i:=0; i<n; i++ {
		f = f.Next
	}
	for f != nil {
		f, s = f.Next, s.Next
	}
	s.Next, s.Next.Next = s.Next.Next, nil
	return newHead.Next
}
