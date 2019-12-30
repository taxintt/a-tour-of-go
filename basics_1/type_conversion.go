package main

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4

	// math.Sqrt(type) - float64
	// uint - 符号なし整数 / int - 符号あり整数
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("%T", f, f)

	var z uint = uint(f)
	fmt.Println(x, y, z)
}