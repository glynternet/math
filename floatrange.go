package math

import "math"

type FloatRange struct {
	From, To float64
}

func (fr FloatRange) Range() float64 {
	return fr.To - fr.From
}

func (fr FloatRange) AbsRange() float64 {
	return math.Abs(fr.Range())
}

func (fr FloatRange) Cap(value float64) float64 {
	max := math.Max(fr.From, fr.To)
	min := math.Min(fr.From, fr.To)
	if value >= max {
		return max
	} else if value <= min {
		return min
	}
	return value
}