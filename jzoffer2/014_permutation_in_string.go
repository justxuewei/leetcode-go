package jzoffer2

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1map := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		if _, ok := s1map[s1[i]]; !ok {
			s1map[s1[i]] = 1
		} else {
			s1map[s1[i]]++
		}
	}

	s2map := make(map[byte]int)

	for i := 0; i < len(s2); i++ {
		if _, ok := s2map[s2[i]]; !ok {
			s2map[s2[i]] = 1
		} else {
			s2map[s2[i]]++
		}

		if i >= len(s1) {
			if v := s2map[s2[i-len(s1)]]; v == 1 {
				delete(s2map, s2[i-len(s1)])
			} else {
				s2map[s2[i-len(s1)]]--
			}
		}

		if len(s1map) == len(s2map) {
			flag := false
			for k, v := range s1map {
				if s2map[k] != v {
					flag = true
					break
				}
			}
			if !flag {
				return true
			}
		}
	}

	return false
}
