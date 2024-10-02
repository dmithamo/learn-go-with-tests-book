package structs

import (
	"math"
)

type Rectangle struct {
	width  float64
	length float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.length)
}

func (r Rectangle) Area() float64 {
	return Rounding(r.width*r.length, 2)
}

func (c Circle) Area() float64 {
	return Rounding(math.Pi*c.radius*c.radius, 2)
}

// Copied from the `net
func Rounding(f float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(f*ratio) / ratio
}
