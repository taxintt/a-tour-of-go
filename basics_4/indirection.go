package main

import (
	"fmt"
)

type Vertex struct{
	X, Y float64
}

func (v *Vertex) Scale (f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64){
	v.X = v.X * f
	v.Y = v.Y * f	
}

func main(){
	v := Vertex{3,4}
	v.Scale(2)
	// pass pointer variable(&v) -> function
	ScaleFunc(&v, 10)

	// make pointer struct(&Vertex)
	p := &Vertex{4,3}
	p.Scale(3)

	// we interpret it (&v).Scale(5) 
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}