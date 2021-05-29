package jzoffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	t.Run("test1", test1)
}

func test1(t *testing.T) {
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian())
	mf.AddNum(3)
	assert.Equal(t, 2.0, mf.FindMedian())
}
