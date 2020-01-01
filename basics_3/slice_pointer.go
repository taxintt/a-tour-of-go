package main

import (
	"fmt"
)

func main() {
	names := [4]string{
		"john",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// names[0] = "xxxx"
	// slices that is came from same array, reflects same changes to slices
	b[0] = "xxxx"
	fmt.Println(a, b)
}