package main

import (
	"fmt"
	"math/rand"
	"time"
)

// we can initialize with initializer outside function 
var p = 4
var i, j int = 1, 2

func main() {
	// Intn->return constant, Seed->return variable
	fmt.Println("My favorite number is", rand.Intn(10))

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Int63())

	// without initializer inside function
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java, p)
}