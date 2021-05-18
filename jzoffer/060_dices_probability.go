package jzoffer

func dicesProbability(n int) []float64 {
	// init tmp
	tmp := make([]float64, 6)
	for i := range tmp {
		tmp[i] = 1.0/6.0
	}
	if n == 1 {
		return tmp
	}
	// dp
	var ret []float64
	for i:=2; i<n+1; i++ {
		ret = make([]float64, 5*i+1)
		for j:=0; j<len(tmp); j++ {
			for k:=0; k<6; k++ {
				ret[j+k] += tmp[j] * 1.0/6.0
			}
		}
		tmp = ret
	}
	return ret
}
