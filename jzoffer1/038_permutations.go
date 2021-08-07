package jzoffer1

import "fmt"

func permutation(s string) []string {
	if s == "" {
		return nil
	} else if len(s) == 1 {
		return []string{s}
	}

	var ret []string

	for i := range s {
		root := s[i]
		p := permutation(s[:i]+s[i+1:])
		for j := range p {
			p[j] = fmt.Sprintf("%c%s", root, p[j])
		}
		ret = append(ret, p...)
	}

	return ret
}
