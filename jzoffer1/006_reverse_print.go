package jzoffer1

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var ret []int

	// reverse linked list
	next1 := head.Next
	head.Next = nil
	var next2 *ListNode
	if next1 != nil {
		next2 = next1.Next
	}

	for {
		if next1 == nil {
			break
		}
		next1.Next = head
		head = next1
		if next2 == nil {
			break
		}
		next1, next2 = next2, next2.Next
	}

	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}
