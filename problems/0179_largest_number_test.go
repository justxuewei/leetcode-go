package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLess(t *testing.T) {
	assert.True(t, less("10", "2"))
	assert.True(t, less("98", "9"))
	assert.True(t, !less("9", "9"))
	assert.True(t, less("34323", "3432"))
	assert.True(t, less("3432", "34329"))
}

func TestQuickSort(t *testing.T) {
	case1 := []string{"3", "30", "34", "5", "9"}
	case1Expect := []string{"9", "5", "34", "3", "30"}
	quickSort(case1)
	assert.Equal(t, case1Expect, case1)

	case2 := []string{"34323", "3432"}
	case2Expect := []string{"3432", "34323"}
	quickSort(case2)
	assert.Equal(t, case2Expect, case2)
}

func TestLargestNumber1(t *testing.T) {
	case1 := []int{0,9,8,7,6,5,4,3,2,1}
	case1Expect := "9876543210"
	assert.Equal(t, case1Expect, largestNumber1(case1))
}
