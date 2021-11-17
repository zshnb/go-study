package strcut

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

const PI = 3.14

func perimeter(rectangle Rectangle) float64 {
	return (rectangle.width + rectangle.height) * 2
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * PI
}
