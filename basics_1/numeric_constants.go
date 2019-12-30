package main

import "fmt"

const (
	// create huge number 
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {return x*10+1}
func needFloat(x float64) float64 {return x * 0.1}

func main(){
	// fmt.Printf("%T", Big)
	// fmt.Printf("%T", Small)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))

	fmt.Println(needFloat(Big))
	fmt.Printf("%T", needFloat(Big))

	// overflows int
	// fmt.Println(needInt(Big))
	// fmt.Printf("%T", needFloat(Big))
}