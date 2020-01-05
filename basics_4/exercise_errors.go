package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64,error){
	// check the value
	if x <0 {
		// If the value is less than 0, return error
		return 0,ErrNegativeSqrt(x)
	}

	z, d := 1.0, 1.0
	for {
		d = (z*z - x) / (2*z)
		z -= d

		if d*d < 1e-10{
			break
		}
	}
	return z, nil
}

// value receiver
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main(){
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	// fmt.Println(ErrNegativeSqrt(-2).Error())
	
}