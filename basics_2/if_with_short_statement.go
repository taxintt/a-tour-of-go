package main

import (
	"fmt"
	"math"
)

// ex) pow(3, 2, 10) -> v=9(=3^2) / lim=10
func pow(x, n, lim float64) float64 {

	// how to write -> if (short statement);(conditions);
	// math.pow：Exponentiation(累乗)
	if v := math.Pow(x, n); v < lim {
		return v
	}else{
		// %g - float64
		fmt.Printf("%g >= %g\n", v, lim)
	}
	
	// can't use v outside if function
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
