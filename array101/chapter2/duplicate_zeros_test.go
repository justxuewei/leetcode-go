package array101

import "testing"

func TestDuplicateZerosCase1(t *testing.T) {
	case1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(case1)
	t.Log(case1)
}

func TestDuplicateZerosCase2(t *testing.T) {
	case2 := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	duplicateZeros(case2)
	t.Log(case2)
}

func TestDuplicateZerosCase3(t *testing.T) {
	case3 := []int{8, 4, 5, 0, 0, 0, 0, 7}
	duplicateZeros(case3)
	t.Log(case3)
}
