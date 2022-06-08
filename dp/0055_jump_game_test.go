package dp

import "testing"

func TestCanJump(t *testing.T) {
	case1 := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	t.Log("case1: ", canJump(case1))
	case2 := []int{3, 2, 1, 0, 4}
	t.Log("case2: ", canJump(case2))
}
