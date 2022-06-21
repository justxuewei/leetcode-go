package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeSquareSum(t *testing.T) {
	assert.True(t, judgeSquareSum(5))
	assert.False(t, judgeSquareSum(3))
	assert.False(t, judgeSquareSum(6))
	assert.True(t, judgeSquareSum(4))
	assert.True(t, judgeSquareSum(2))
}
