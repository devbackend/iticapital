package money

import "math"

type Money float64

func (p *Money) Val() Money {
	return Money(math.Round(float64(*p)*10000) / 10000)
}
