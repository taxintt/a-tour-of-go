package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	// fixed value
	// for i:=0; i<10; i++ {
	// 	z -= (z*z - x) / (2*z)
	// 	fmt.Println(z)
	// }

	// minimize difference (starts from d(=1.0))
	// for d:=1.0; d*d > 1e-10; z -= d {
	// 	d = (z*z - x) / (2*z)
	// }

	// my answer
	d := 1.0
	for {
		d = (z*z - x) / (2*z)
		z -= d

		if d*d < 1e-10{
			break
		}
	}
	return z
}

func main() {
	// similar result
	var i = 5.0

	fmt.Println(Sqrt(i))
	fmt.Println(math.Sqrt(i))
}