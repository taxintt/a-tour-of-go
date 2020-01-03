package main

import (
	"fmt"
	"math"
)

// type
type Vertex struct {
	x, y float64
}

// bind method to type(Vertex)
// メソッドは、レシーバ引数(v Vertex)を伴う関数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func normalAbs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3,4}

	fmt.Println(v.Abs())
	fmt.Println(normalAbs(v))
} 