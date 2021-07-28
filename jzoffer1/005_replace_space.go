package jzoffer1

import "strings"

func replaceSpace(s string) string {
	var sb strings.Builder

	for i := range s {
		if s[i] == ' ' {
			sb.WriteString("%20")
			continue
		}
		sb.WriteByte(s[i])
	}

	return sb.String()
}
