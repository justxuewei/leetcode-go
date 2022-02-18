package jzoffer2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAtoiInBinary(t *testing.T) {
	assert.Equal(t, 2, atoiInBinary("10"))
}

func TestItoAInBinary(t *testing.T) {
	assert.Equal(t, "10", itoaInBinary(2))
}

func TestAddBinary(t *testing.T) {
	assert.Equal(t, "100", addBinary("11", "1"))
}
