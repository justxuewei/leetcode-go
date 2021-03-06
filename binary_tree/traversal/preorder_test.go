package traversal

import (
	btree "github.com/xavier-niu/leetcode/binary_tree"
	"testing"
)

func TestPreorderTraversalRecursively(t *testing.T) {
	case1, _ := btree.GenerateBinaryTree([]int{1, btree.Null, 2, 3})
	output1 := preorderTraversalRecursively(case1, nil)
	t.Log(output1)
}

func TestPreorderTraversalIteratively(t *testing.T) {
	case1, _ := btree.GenerateBinaryTree([]int{1, btree.Null, 2, 3})
	output1 := preorderTraversalIteratively(case1)
	t.Log(output1)
}
