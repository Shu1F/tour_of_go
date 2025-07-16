package main

import (
	"fmt"
	"math"
)

type Vertex_methods struct {
	X, Y float64
}

func (v Vertex_methods) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vertex_methods{3, 4}
	fmt.Println(v.Abs())
}
