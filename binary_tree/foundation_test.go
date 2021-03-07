package foundation

import "testing"

func TestGenerateBinaryTree(t *testing.T) {
	case1 := []int{1, Null, 2, 3}
	case1Tree, err := GenerateBinaryTree(case1)
	if err != nil {
		t.Error(err)
	}
	t.Log(case1Tree)

	var case2 []int
	case2Tree, err := GenerateBinaryTree(case2)
	if err != nil {
		t.Error(err)
	}
	t.Log(case2Tree)

	case3 := []int{1}
	case3Tree, err := GenerateBinaryTree(case3)
	if err != nil {
		t.Error(err)
	}
	t.Log(case3Tree)
}
