package main

import (
	"fmt"
	"math"
)

type Vertex_methods_pointer struct {
	X, Y float64
}

func (v Vertex_methods_pointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex_methods_pointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex_methods_pointer{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
