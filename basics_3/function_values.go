package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64)float64) float64{
	// hypot(3, 4)
	return fn(3, 4)
}

func hypot(x, y float64) float64{
	return math.Sqrt(x*x + y*y)
}

func main(){
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))

	// math.Pow: Exponentiation - 累乗(3^4)
	fmt.Println(compute(math.Pow))
}