package main

import (
	"fmt"
)

func main(){
	// slice(nil)
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil{
		fmt.Println("nil!")
	}
}