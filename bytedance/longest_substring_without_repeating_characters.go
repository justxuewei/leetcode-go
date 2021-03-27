package bytedance

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 { return 0 }
	longestLen := 0
	ws, we := 0, 0
	for i:=1; i<len(s); i++ {
		// search duplicated elements
		for j:=ws; j<=we; j++ {
			if s[j] == s[i] {
				l := we - ws + 1
				fmt.Println(l)
				if longestLen < l {
					longestLen = l
				}
				ws = j + 1
				break
			}
		}
		we = i
	}
	lastLen := we - ws + 1
	fmt.Println(lastLen)
	if lastLen > longestLen {
		longestLen = lastLen
	}
	return longestLen
}
