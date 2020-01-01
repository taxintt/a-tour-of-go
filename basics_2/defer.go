package main

import (
	"fmt"
)

func main(){
	// postpone call defer function after return statement (ex. disconnect DB connection)
	// defer - LIFO(Last in First out)
	defer deferfunc()

	// 1
	fmt.Println("hello world")	
}

func deferfunc(){
	// Anonymous function - func(){...}()
	defer func(){
		// 3
		fmt.Println("Finish func")
	}()
	// 2
	fmt.Println("Begin func")
}