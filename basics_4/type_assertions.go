package main

import (
	"fmt"
)

func main(){
	var i interface{} = "hello"

	// t := i.(T)
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// doesn't happen panic (because of ok(bool) variable)
	f, ok := i.(float64) 
	fmt.Println(f, ok)

	// panic (because of type's disagreement)
	f = i.(float64) 
	fmt.Println(f)
}