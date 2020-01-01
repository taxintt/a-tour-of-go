package main

import (
	"fmt"
)

func main(){
	// making slice -- make(slice, length, capacity)
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	// expand length 
	c := b[:2]
	printSlice("c", c)

	// reduct the length
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v \n",
		s, len(x), cap(x), x)
}