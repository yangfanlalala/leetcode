package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}
	k = k % l
	if k == 0 {
		return head
	}
	cur = head
	for i := 0; i < l-k-1; i++ {
		cur = cur.Next
	}
	cur.Next, head, cur = nil, cur.Next, head
	c := head
	for {
		if c.Next == nil {
			c.Next = cur
			break
		}
		c = c.Next
	}
	return head
}
