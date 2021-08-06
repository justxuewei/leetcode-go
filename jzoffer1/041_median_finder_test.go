package jzoffer1

import "testing"

func TestMedianFinder(t *testing.T) {
	finder := NewMedianFinder()
	finder.AddNum(-1)
	finder.AddNum(-2)
	finder.AddNum(-3)
	t.Log(finder.FindMedian())
}
