package bytedance

import "testing"

func TestHeapSort(t *testing.T) {
	case1 := []int{5, 2, 3, 1}
	heapSort(case1)
	t.Log(case1)

	case2 := []int{-2, 3, -5}
	heapSort(case2)
	t.Log(case2)

	case3 := []int{5, 2, 6}
	heapSort(case3)
	t.Log(case3)
}
