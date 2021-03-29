package bytedance

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	case1Head := &ListNode{Val: 4}
	case1P1 := &ListNode{Val: 2}
	case1P2 := &ListNode{Val: 1}
	case1P3 := &ListNode{Val: 3}
	case1Head.Next = case1P1
	case1P1.Next = case1P2
	case1P2.Next = case1P3
	case1NewHead := sortList(case1Head)
	case1P := case1NewHead
	for case1P != nil {
		fmt.Println(case1P.Val)
		case1P = case1P.Next
	}
}
