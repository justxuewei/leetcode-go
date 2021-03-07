package btree

import (
	"testing"
)

func TestPreorderTraversalRecursively(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{1, Null, 2, 3})
	output1 := preorderTraversalRecursively(case1, nil)
	t.Log(output1)
}

func TestPreorderTraversalIteratively(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{1, Null, 2, 3})
	output1 := preorderTraversalIteratively(case1)
	t.Log(output1)
}
