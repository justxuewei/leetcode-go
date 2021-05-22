package jzoffer

func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	// init array b
	b := make([]int, len(a))
	for i := range b {
		b[i] = 1
	}
	// lower triangle
	for i:=1; i<len(b); i++ {
		b[i] = a[i-1] * b[i-1]
	}
	// upper triangle
	tmp := 1
	for i:=len(b)-2; i>=0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
