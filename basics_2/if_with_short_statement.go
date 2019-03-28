package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//math.pow：累乗
	//pow(3, 2, 10) -> 9(3^2)<10:T - 9, 27(3^3)<20:F - 20
	if v := math.Pow(x, n); v < lim {
		return v
	}else{
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
