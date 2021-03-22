package bytedance

import "testing"

func TestSubset(t *testing.T) {
	case1 := []int{9, 0, 3, 5, 7}
	case1Ret := subsets(case1)
	t.Log(case1Ret)
}
