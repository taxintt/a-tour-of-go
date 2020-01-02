package main

import (
	"fmt"
)

func main(){
	pow := make([]int, 10)

	fmt.Println("first loop...")
	for i := range pow {
		fmt.Println(i)
		pow[i] = 1 << uint(i) // 2**1
	}

	// only use value (not use index)
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}