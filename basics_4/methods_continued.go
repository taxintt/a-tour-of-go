package main

import (
	"fmt"
	"math"
)

type Myfloat float64 

// method
func (f Myfloat) Abs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

func main(){
	f := Myfloat(-math.Sqrt2) 
	// f := Myfloat(math.Sqrt(3)) is OK

	fmt.Println(f.Abs())
}