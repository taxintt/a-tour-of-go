package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		fmt.Println(x)
		return sqrt(-x) + "i"
	}
	// return formatted string（fmt.Sprint）
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	
	// if we use variables, you can set type(float64) 
	var i = 2.0
	fmt.Println(sqrt(i), sqrt(-4))
}
