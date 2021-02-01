package chapter6

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

	case4 := []int{5, 1, 2, 3, 4}
	quickSort(case4)
	t.Log(case4)
}

func TestHeightChecker(t *testing.T) {

	case1 := []int{5, 1, 2, 3, 4}
	t.Log(heightChecker(case1))

}
