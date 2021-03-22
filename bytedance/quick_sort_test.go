package bytedance

import "testing"

func TestQuickSort(t *testing.T) {
	case1 := []int{2, 1, 3}
	quickSort(case1)
	t.Log(case1)

	case2 := []int{3, 2, 1}
	quickSort(case2)
	t.Log(case2)

	case3 := []int{1, 2, 3}
	quickSort(case3)
	t.Log(case3)

	case4 := []int{4, 3, 1, 6, 2, 4}
	quickSort(case4)
	t.Log(case4)

	case5 := []int{1, 2, 5, 4, 3, -1}
	quickSort(case5)
	t.Log(case5)
}
