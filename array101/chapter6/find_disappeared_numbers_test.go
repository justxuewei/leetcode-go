package chapter6

import "testing"

func TestFindDisappearedNumbers(t *testing.T) {
	case1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	t.Log(findDisappearedNumbers(case1))
}
