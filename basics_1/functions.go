package main

import "fmt"

// can omit type of variable (ex. add(x, y int))
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}