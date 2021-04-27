package jzoffer

import "testing"

func TestIsNumber(t *testing.T) {
	case1 := " +100.2e-4 "
	if !isNumber(case1) {
		t.Errorf("case1 is a valid number")
	}
}
