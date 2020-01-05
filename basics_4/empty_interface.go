package main

import (
	"fmt"
)

func main(){
	var i interface{}
	describe(i)

	// can hold values of any types
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}){
	fmt.Printf("(%v, %T)\n", i, i)
}