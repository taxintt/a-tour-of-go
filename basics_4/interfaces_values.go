package main

import (
	"fmt"
	"math"
)

// pattern - 1.pointer receiver
type I interface{
	M()
}

type T struct{
	S string
}

func (t *T) M(){
	fmt.Println(t.S)
}

// pattern - 2.value receiver
type F float64

func (f F) M(){
	fmt.Println(f)
}

func main(){
	var i I

	// 1.pointer receiver
	i = &T{"Hello"} // address
	describe(i)
	i.M()

	// 2.value receiver
	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

