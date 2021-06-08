package main

import (
	"fmt"

	"github.com/divjotarora/go-interfaces/calculator"
	"github.com/divjotarora/go-interfaces/shapes"
)

type mockShape struct {
}

func (m *mockShape) SideLengths() []int32 {
	return nil
}

func main() {
	calc := calculator.NewCalculator()

	// Use known types
	printPerimeter(calc, shapes.NewSquare(5))
	printPerimeter(calc, shapes.NewRectangle(4, 5))

	// Use mock type to test error case
	printPerimeter(calc, &mockShape{})
}

func printPerimeter(calc *calculator.Calculator, shape calculator.Shape) {
	p, err := calc.Perimeter(shape)
	if err != nil {
		fmt.Printf("err calculating perimeter: %v\n", err)
		return
	}
	fmt.Printf("perimeter: %d\n", p)
}
