package main

import "fmt"

//{}の外ではvarのみで宣言する
var i, j int = 1, 2

func main() {
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	fmt.Println("My favorite number is", rand.Intn(10))
}