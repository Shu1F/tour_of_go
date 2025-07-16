package main

import "fmt"

type Vertex4 struct {
	X, Y int
}

var (
	v1 = Vertex4{1, 2}
	v2 = Vertex4{X: 2}
	v3 = Vertex4{}
	p = &Vertex4{1, 2}
	q = &v2

)

func main() {
	fmt.Println(v1, v2, v3, p, q)
}
