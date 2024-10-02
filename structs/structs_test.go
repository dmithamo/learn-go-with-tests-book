package structs

import "testing"

func TestStruct(t *testing.T) {
	rectGiven := Rectangle{
		width:  20.04,
		length: 30.40,
	}
	circlGiven := Circle{10.0}

	t.Run("rect area", func(t *testing.T) {
		want := 609.22
		checkArea(t, rectGiven, want)
	})

	t.Run("circle area", func(t *testing.T) {
		want := 314.16
		checkArea(t, circlGiven, want)
	})
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAreaTableDriven(t *testing.T) {
	tests := []struct {
		shape Shape
		want  float64
	}{
		{Circle{10.0}, 314.159},
		{Rectangle{10.0, 20.0}, 200.00},
	}

	for _, test := range tests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("%#v got %v want %v", test.shape, got, test.want)
		}
	}
}
