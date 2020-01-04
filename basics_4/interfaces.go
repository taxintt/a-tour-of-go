package main

import (
	"fmt"
	"math"
)

// declare interface
type Abser interface{
	Abs() float64
}

type Vertex struct{
	X, Y float64
}

// pointer receiver -> Be considered that it does not implement Abs()
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// value receiver -> implement Abs()
func (f MyFloat) Abs() float64{
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

func main(){

	// define variable by using interface
	var a Abser

	f := MyFloat(-math.Sqrt2)
	a = f // a MyFloat implements Abser

	v := Vertex{3,4}
	a = &v // [work]
	a = v // [doesn't work] -> a *Vertex(Not Vertex) implements Abser

	fmt.Println(a.Abs())
}

