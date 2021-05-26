package jzoffer

import (
	btree "github.com/xavier-niu/leetcode/binary_tree"
	"testing"
)

func TestSerialize(t *testing.T) {
	btreeCodec := NewBTreeCodec()
	t.Logf("root=nil: %s", btreeCodec.serialize(nil))
	case1, _ := btree.GenerateBinaryTree([]int{1,2,3,btree.Null,btree.Null,4,5})
	t.Logf("case1: %s", btreeCodec.serialize(case1))
}
