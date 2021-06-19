package btree

import "testing"

func TestHasPathSum(t *testing.T) {
	case1, _ := GenerateBinaryTree([]int{1, 2})
	t.Log(hasPathSum(case1, 1))
}
