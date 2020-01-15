package main

import (
	"fmt"
	"time"
)

func main(){
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(100 * time.Millisecond)

	for{
		select{
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			// receiving from c would block
			fmt.Println("  .")
			time.Sleep(50 * time.Millisecond)
		}
		
	}
}