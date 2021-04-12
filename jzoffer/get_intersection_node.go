package jzoffer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { return nil }

	pA, pB := headA, headB

	lenA, lenB := 0, 0
	for pA != nil {
		pA = pA.Next
		lenA++
	}
	for pB != nil {
		pB = pB.Next
		lenB++
	}

	pA, pB = headA, headB
	if lenA < lenB {
		remaining := lenB - lenA
		for remaining > 0 {
			pB = pB.Next
			remaining--
		}
	} else if lenA > lenB {
		remaining := lenA - lenB
		for remaining > 0 {
			pA = pA.Next
			remaining--
		}
	}

	for pA != nil && pA != pB {
		pA, pB = pA.Next, pB.Next
	}

	return pA
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { return nil }

	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
