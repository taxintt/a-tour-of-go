package main

import (
	"fmt"
)

// interface -> type of variables is specified
func do(i interface{}){
	switch v := i.(type){

	// specify type of variables
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main(){
	do(21)
	do("21")
	do(true)
}