package btree

import "testing"

func TestMaxDepth(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{3, 9, 20, Null, Null, 15, 7})
	output1 := maxDepth(case1)
	t.Log(output1)
}
