package golang_united_school_homework

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c *Circle) CalcPerimeter() float64 {
	pi := 3.14
	return 2 * pi * c.Radius
}

func (c *Circle) CalcArea() float64 {
	pi := 3.14
	return pi * (c.Radius * c.Radius)
}
