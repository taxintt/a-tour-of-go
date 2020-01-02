package main

import (
	"fmt"
)

// we need var(for definition) because the variable is out of function(main)
var pow = []int{1,2,4,8,16,32,64,128}

func main(){
	// Golang
	// for <index>, <value> := range <array> 
	for i, v := range pow{
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Python
	// for x in range(len(pow)):
	// print(x)
	// print(pow[x])
}