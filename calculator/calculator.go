package calculator

import "errors"

// We're the consumer, so we define the interface. It's not defined in the "shapes" package
type Shape interface {
	SideLengths() []int32
}

type Calculator struct {
}

// Return structs
func NewCalculator() *Calculator {
	return &Calculator{}
}

// Accept interfaces - we accept any type that fulfills the "Shape" interface
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
