package main

import (
	"fmt"
)

type Vertex struct{
	x int
	y int
}

func main(){
	v := Vertex{1, 2}
	p := &v // set pointer to v

	fmt.Println(p) 
	fmt.Println(*p) // read value from the pointer

	p.x = 1e9 // (*p).x is also ok
	fmt.Println(v)
}