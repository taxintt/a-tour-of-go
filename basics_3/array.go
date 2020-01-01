package main

import (
	"fmt"
)

func main(){
	var arr [2]string // declare array(fixed length)

	arr[0] = "hello"
	arr[1] = "world"
	fmt.Println(arr)

	primes := [6]int{1,3,5,7,9,11}
	sum:= 0
	for i := 0; i<len(primes); i++ {
		sum += primes[i]
	}
	fmt.Println(sum)
}