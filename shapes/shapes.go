package shapes

// Everything here returns *structs*

type Square struct {
	sideLength int32
}

func NewSquare(sideLength int32) *Square {
	return &Square{sideLength}
}

func (s *Square) SideLengths() []int32 {
	return []int32{s.sideLength, s.sideLength, s.sideLength, s.sideLength}
}

type Rectangle struct {
	x, y int32
}

func NewRectangle(x, y int32) *Rectangle {
	return &Rectangle{x, y}
}

func (r *Rectangle) SideLengths() []int32 {
	return []int32{r.x, r.x, r.y, r.y}
}
