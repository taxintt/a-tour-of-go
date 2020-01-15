package main

import (
	"fmt"
)

func main(){
	// buffer=1 -> fatal error: all goroutines are asleep - deadlock!
	// 指定しないと容量は0
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}