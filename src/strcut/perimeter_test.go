package strcut

import "testing"

func TestPerimeter(t *testing.T) {
	actual := perimeter(Rectangle{10, 10})
	expected := 40.0
	if actual != expected {
		t.Errorf("expected '%.2f' but got '%.2f'", expected, actual)
	}
}

func checkArea(t *testing.T, shape Shape, expected float64) {
	if shape.area() != expected {
		t.Errorf("expected '%.2f' but got '%.2f'", expected, shape.area())
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle's area", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		expected := 100.0
		checkArea(t, rectangle, expected)
	})
	t.Run("circle's area", func(t *testing.T) {
		circle := Circle{1}
		expected := 3.14
		checkArea(t, circle, expected)
	})
}
