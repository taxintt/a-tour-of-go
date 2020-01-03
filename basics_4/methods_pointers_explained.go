package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

// function
func Abs(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function - need * (for assign value)
func Scale(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{3,4}

	// pass pointer variable(&v) -> function
	Scale(&v, 10)
	fmt.Println(Abs(v))
}