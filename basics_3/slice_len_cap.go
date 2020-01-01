package main

import (
	"fmt"
)

func main(){
	s := []int{2,3,5,7,11,13}
	printSlice(s)

	// give slice to zero length(->len)
	s = s[:0]
	printSlice(s)	
	
	// extend its length(->len)
	s = s[:4]
	printSlice(s)	

	// drop first two values(->cap)
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int){
	// %d(verb) - unsigned integer type, %v(verb) - display values
	// len(s) - length of array
	// cap(s) - number of elements that is origin
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}