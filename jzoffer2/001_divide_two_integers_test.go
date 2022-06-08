package jzoffer2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideTwoIntegers(t *testing.T) {

	t.Run("case1", func(t *testing.T) {
		assert.Equal(t, 7, divide(15, 2))
		assert.Equal(t, -7, divide(-15, 2))
		assert.Equal(t, 1, divide(2, 2))
		assert.Equal(t, -1073741824, divide(-2147483648, 2))
	})

}
