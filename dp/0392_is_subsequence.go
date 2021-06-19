package dp

func isSubsequence(s string, t string) bool {
	if len(s) == 0 { return true }
	if len(t) == 0 { return false }
	matchidx := 0
	for i := range t {
		if t[i] == s[matchidx] {
			matchidx++
		}
		if matchidx == len(s) {
			return true
		}
	}
	return false
}
