package jzoffer2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationInString(t *testing.T) {
	assert.True(t, checkInclusion("ab", "eidbaooo"))
}
