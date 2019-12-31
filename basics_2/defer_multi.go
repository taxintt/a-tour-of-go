package main

import (
	"fmt"
)

func main(){
	fmt.Println("counting")

	for i:=0; i<10; i++ {
		// defer - LIFO(Last in First out)
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

