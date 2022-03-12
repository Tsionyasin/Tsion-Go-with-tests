package smi

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
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

	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	want := 72.00

	if got != want {

		t.Errorf("got %.2f, want %.2f", got, want)

	}

}

func Parameter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Width + rectangle.Height)

}

func Area(rectangle Rectangle) float64 {

	return rectangle.Width * rectangle.Height
}
