package lnklist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	var p1 *ListNode
	ret, p2, p3 := head, head, head
	remaining := n

	for remaining > 0 {
		p3 = p3.Next
		remaining--
	}

	for p3 != nil {
		p1 = p2
		p2, p3 = p2.Next, p3.Next
	}

	// goland:noinspection GoNilness
	if p1 != nil {
		p1.Next = p2.Next
	} else {
		ret = head.Next
	}

	return ret
}
