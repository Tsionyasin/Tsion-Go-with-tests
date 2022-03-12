package smi

import "testing"
import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func TestParameter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Parameter(rectangle)
	want := 40.00

	if got != want {

		t.Errorf("got %.2f, want %.2f", got, want)

	}

}
func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func Parameter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Width + rectangle.Height)

}

func Area(rectangle Rectangle) float64 {

	return rectangle.Width * rectangle.Height
}
