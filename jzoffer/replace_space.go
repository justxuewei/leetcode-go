package jzoffer

import "strings"

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	ret := ""
	for _, v := range s {
		str := string(v)
		if str == " " {
			ret += "%20"
		} else {
			ret += str
		}
	}
	return ret
}

func replaceSpace1(s string) string {
	var ret strings.Builder
	for i := range s {
		if s[i] == ' ' {
			ret.WriteString("%20")
		} else {
			ret.WriteByte(s[i])
		}
	}
	return ret.String()
}
