package array101

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	case1 := []int{0}

	if length := findMaxConsecutiveOnes(case1); length != 0 {
		t.Error(length)
	}
}
