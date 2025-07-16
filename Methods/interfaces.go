package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs()float64
}

func main() {
	var a Abser
	f := MyFloat_interface(-math.Sqrt2)
	v := Vertex_interface{3, 4}

	a = f
	a = &v

	a = v

	fmt.Println(a.Abs())
}

type MyFloat_interface float64

func (f MyFloat_interface) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex_interface struct {
	X, Y float64
}

func (v * Vertex_interface) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
