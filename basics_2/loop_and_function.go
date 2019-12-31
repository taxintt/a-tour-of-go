package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i<3; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	// similar result
	var i = 2.0

	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}