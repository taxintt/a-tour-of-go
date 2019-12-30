package main

import "fmt"

const Pi float32 = 3.14

func main(){
	// we can't use := about const 
	// world(variable) / World(constants)
	const World = "世界"
	fmt.Println("Hello", World)

	fmt.Printf("Happy %T", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}