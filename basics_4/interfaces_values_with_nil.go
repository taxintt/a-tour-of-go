package main

import (
	"fmt"
)

type I interface{
	M()
}

type T struct{
	S string
}

func (t *T) M(){
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main(){
	var i I

	var t *T
	// var t T -> doesn't work ()
	i = t
	describe(i) // (<nil>, *main.T) - doesn't be assigned values to i
	i.M()

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}