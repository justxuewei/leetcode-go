package jzoffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePairs(t *testing.T) {
	case1 := []int{7, 5, 6, 4}
	assert.Equal(t, reversePairs(case1), 5)
}
