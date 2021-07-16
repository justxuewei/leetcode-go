package dp

import "testing"

func TestIsInterleave(t *testing.T) {
	case1 := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	t.Log(case1)
}
