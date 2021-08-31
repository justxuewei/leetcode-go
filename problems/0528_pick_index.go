package problems

import (
	"math/rand"
	"sort"
)

type WeightedPicker struct {
	weights []int
}

func NewWeightedPicker(w []int) WeightedPicker {
	for i:=1; i<len(w); i++ {
		w[i] += w[i-1]
	}
	return WeightedPicker{
		weights: w,
	}
}

func (p *WeightedPicker) PickIndex() int {
	// generate numbers in [1, p.weights[len(p.weights)-1]]
	index := rand.Intn(p.weights[len(p.weights)-1]) + 1
	return sort.SearchInts(p.weights, index)
}
