package main

import (
	"fmt"
)


type I interface{
	M()
}

func main(){
	var i I 
	describe(i)

	// panic: runtime error: 
	// invalid memory address or 【nil pointer dereference】
	i.M()
}

func describe(i I){
	fmt.Printf("(%v, %T)\n",i,i)
}