package dp

import "testing"

func TestMinPathSum(t *testing.T) {
	case1 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	t.Log(minPathSum(case1))
}
