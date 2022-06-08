package bytedance

import (
	"testing"

	btree "github.com/xavier-niu/leetcode/btree"
)

func TestPartSumBytedance(t *testing.T) {
	case1 := []int{5, 4, 8, 11, btree.Null, 13, 4, 7, 2, btree.Null, btree.Null, btree.Null, 5, 1}
	case1Root, _ := btree.GenerateBinaryTree(case1)
	t.Log(partSumBytedance(case1Root, 22))
}
