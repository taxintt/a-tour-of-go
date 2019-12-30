package main

import "fmt"

//  declare type of return
func swap(x, y string) (string, string) {
	return y, x
}

// := declaration and assignment
// we can only use := in functions
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}