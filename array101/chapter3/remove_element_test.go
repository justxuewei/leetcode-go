package chapter3

import "testing"

func TestRemoveElement1(t *testing.T) {
	case1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	length := removeElement(case1, 2)
	for i := 0; i < length; i++ {
		t.Log(case1[i])
	}
}
