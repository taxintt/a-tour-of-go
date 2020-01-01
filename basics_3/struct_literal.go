package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // x is specified and y:0 is implicit
	v3 = Vertex{} // X:0 and Y:0 is implicit
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}