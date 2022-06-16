package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoordination(t *testing.T) {
	case1Row, case1Col := getCoordination(4, 1)
	case1ExpectRow, case1ExpectCol := 3, 0
	assert.Equal(t, case1ExpectRow, case1Row)
	assert.Equal(t, case1ExpectCol, case1Col)

	case2Row, case2Col := getCoordination(4, 16)
	case2ExpectRow, case2ExpectCol := 0, 0
	assert.Equal(t, case2ExpectRow, case2Row)
	assert.Equal(t, case2ExpectCol, case2Col)

	case3Row, case3Col := getCoordination(4, 12)
	case3ExpectRow, case3ExpectCol := 1, 3
	assert.Equal(t, case3ExpectRow, case3Row)
	assert.Equal(t, case3ExpectCol, case3Col)

	case4Row, case4Col := getCoordination(4, 6)
	case4ExpectRow, case4ExpectCol := 2, 2
	assert.Equal(t, case4ExpectRow, case4Row)
	assert.Equal(t, case4ExpectCol, case4Col)
}

func TestSnakesAndLadders(t *testing.T) {
	case1 := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}
	assert.Equal(t, 4, snakesAndLadders(case1))

	case2 := [][]int{
		{-1, 4, -1},
		{6, 2, 6},
		{-1, 3, -1},
	}
	assert.Equal(t, 2, snakesAndLadders(case2))
}
