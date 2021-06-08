package calculator

import "errors"

type Shape interface {
	SideLengths() []int32
}

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Perimeter(s Shape) (int32, error) {
	sides := s.SideLengths()
	if len(sides) == 0 {
		return 0, errors.New("shape cannot have 0 sides")
	}

	var sum int32
	for _, side := range sides {
		sum += side
	}
	return sum, nil
}
