package btree

import "testing"

func TestLevelOrder(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{3, 9, 20, Null, Null, 15, 7})
	output1 := levelOrder(case1)
	t.Log(output1)
}
