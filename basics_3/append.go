package main

import (
	"fmt"
)

func main(){
	var s []int
	printSlice(s)

	// append(slice,element)
	s = append(s,0)
	printSlice(s)

	s = append(s,1,2) // +1
	printSlice(s)

	// if the capacity is shorter at the time we add variables, allocate bigger array
	// return value of slice indicates new place of assignment
	s = append(s,3,4,5) // +1
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d, cap=%d %v \n", len(s), cap(s), s)
}