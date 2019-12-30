package main

import (
	"fmt"
	"math/rand"
)

func main(){
	// rand.Intn: return constant
	fmt.Println("My favorite number is", rand.Intn(10))
}