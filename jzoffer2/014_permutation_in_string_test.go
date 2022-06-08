package jzoffer2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutationInString(t *testing.T) {
	assert.True(t, checkInclusion("ab", "eidbaooo"))
}
