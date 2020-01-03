package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

// 1. method - value receiver
func (v Vertex) Abs() float64{	
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func normalAbs(v Vertex) float64 {
// 	return math.Sqrt(v.x*v.x + v.y*v.y)
// }

// 2. method - pointer receiver
// メソッド内で構造体の値を書き換える必要がある場合に利用する 
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{3,4}

	v.Scale(10)
	fmt.Println(v.Abs())
}