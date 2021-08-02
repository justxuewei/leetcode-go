package jzoffer1

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for ; k>1 && fast!=nil; k-- {
		fast = fast.Next
	}
	if k != 1 {
		return head
	}

	for fast.Next != nil {
		head, fast = head.Next, fast.Next
	}

	return head
}
