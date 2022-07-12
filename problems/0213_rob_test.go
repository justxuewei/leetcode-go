package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type robTestCase struct {
	name     string
	input    []int
	expected int
}

func TestRob(t *testing.T) {
	tests := []robTestCase{
		{"testCase1", []int{1, 3, 1, 3, 100}, 103},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, rob(test.input))
		})
	}
}
