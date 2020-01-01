package main

import (
	"fmt"
)

func main(){
	// if we change [values that is used between functions], we should use variables
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(p) // p(pointer) - address of variable
	fmt.Println(*p) // read i through the pointer(p)

	*p = 21 // set i through the pointer
	fmt.Println(&i)
	fmt.Println(i) // see the new value of i(21)

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j(73)
}