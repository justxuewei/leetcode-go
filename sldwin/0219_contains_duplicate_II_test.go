package sldwin

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	t.Log(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
