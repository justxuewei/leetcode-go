package dp

import "testing"

func TestIsInterleave(t *testing.T) {
	t.Log(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	t.Log(isInterleave("", "", "a"))
}
