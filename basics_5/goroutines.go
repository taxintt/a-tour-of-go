package main

import (
	"fmt"
	// "time"
)

// another goroutine - say
func say(s string) {
	for i := 0; i<5; i++ {
		// sleepを無くすとhello(main) -> world(say goroutine)の順で実行される
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// main goroutine
func main(){
	go say("world")
	say("hello")
}

// Output
// hello, worldを5回ずつ実行 - それぞれの順番は保障されない