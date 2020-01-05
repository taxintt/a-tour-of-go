package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// method receiver
// String() - method that is inherited by interface(stringer)
func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error{
	// return value(&MyError)
	return &MyError{
		time.Now(),
		"it doesn't work...",
	}
}

func main(){
	if err := run(); err != nil{
		fmt.Println(err)
	}
}