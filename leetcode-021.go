package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val > l2.Val {
		head, l2 = l2, l2.Next
	} else {
		head, l1 = l1, l1.Next
	}
	c := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			c.Next, c, l2 = l2, l2, l2.Next
		} else {
			c.Next, c, l1 = l1, l1, l1.Next
		}
	}
	if l1 != nil {
		c.Next = l1
	} else if l2!=nil {
		c.Next = l2
	}
	return head
}
