package btree

import "testing"

func TestPostorderTraversalIteratively(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{1, Null, 2, 3})
	output1 := postorderTraversalIteratively(case1)
	t.Log(output1)
}
