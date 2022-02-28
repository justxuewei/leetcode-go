package jzoffer2

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	s1map := make(map[int32]int)

	for _, b := range s1 {
		if _, ok := s1map[b]; !ok {
			s1map[b] = 1
		} else {
			s1map[b]++
		}
	}

	s2map := make(map[int32]int)

	for i := 0; i < len(s2); i++ {

	}

	return true
}
