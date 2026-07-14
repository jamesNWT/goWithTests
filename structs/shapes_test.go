package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t *testing.T) {

	type areaTest struct {
		name    string
		shape   Shape
		hasArea float64
	}

	areaTests := []areaTest{
		{name: "Rectangle", shape: Rectangle{Width: 5.0, Height: 6.0}, hasArea: 30.0},
		{name: "Circle", shape: Circle{Radius: 1.0}, hasArea: math.Pi},
		{name: "Triangle", shape: Triangle{Base: 6.0, Height: 12.0}, hasArea: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			want := test.hasArea

			if got != want {
				t.Errorf("%#v, got %.2f, want %.2f", test, got, want)
			}
		})
	}
}
