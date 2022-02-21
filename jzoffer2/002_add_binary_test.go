package jzoffer2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringCompletion(t *testing.T) {
	var a, b string

	a, b = stringCompletion("11", "1")
	assert.Equal(t, "11", a)
	assert.Equal(t, "01", b)

	a, b = stringCompletion("1", "11")
	assert.Equal(t, "11", a)
	assert.Equal(t, "01", b)
}

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "100", addBinary("1", "11"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
}
