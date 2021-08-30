package problems

import "time"

type WeightedPicker struct {
	weights []int
	length  int
	factor  int
}

func NewWeightedPicker(w []int) WeightedPicker {
	var length int
	for _, v := range w {
		length += v
	}
	return WeightedPicker{
		weights: w,
		length:  length,
		factor:  3,
	}
}

func (p *WeightedPicker) PickIndex() int {
	rand := int(int64(p.factor) * time.Now().Unix() % int64(p.length))
	for i, v := range p.weights {
		if rand = rand - v; rand < 0 {
			return i
		}
	}
	return 0
}
