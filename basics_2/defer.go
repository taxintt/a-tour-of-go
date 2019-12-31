package main

import (
	"fmt"
)

func main(){
	// postpone call defer function after return statement (ex. disconnect DB connection)
	// defer - LIFO(Last in First out)
	defer deferfunc()

	fmt.Println("hello world")	
}

func deferfunc(){
	defer func(){
		fmt.Println("Finish func")
	}()
	fmt.Println("Begin func")
}