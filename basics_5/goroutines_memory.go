package main

import (
	"fmt"
	"time"
)

// goroutine - 並行処理(Concurrency):同時に質の異なることを行う
func main()  {
    fmt.Println("main start")
    for i := 0; i < 5; i++ {
        go func() {
			// mainのgoroutineでのiの状態を参照する(処理の実行によってiが変わる)
            Greet(i)
		}()
    }
    time.Sleep(time.Second)
    fmt.Println("main end")
}

func Greet(i int) {
    fmt.Println("hello",i)
}