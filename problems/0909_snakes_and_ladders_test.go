package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNum(t *testing.T) {
	assert.Equal(t, 1, getNum(4, 3, 0))
	assert.Equal(t, 4, getNum(4, 3, 3))
	assert.Equal(t, 16, getNum(4, 0, 0))
	assert.Equal(t, 10, getNum(4, 1, 1))
}

func TestGetCoordination(t *testing.T) {
	case1Row, case1Col := getCoordination(4, 1)
	case1ExpectRow, case1ExpectCol := 3, 0
	assert.Equal(t, case1ExpectRow, case1Row)
	assert.Equal(t, case1ExpectCol, case1Col)
}
