package main

import (
	"fmt"
	"strconv"
)

func main(){
	sum := 0
	for i:=0;  i<10;  i++ {
		sum += i
		fmt.Println(i)
		fmt.Println("first loop: " + strconv.Itoa(sum))
	}

	// can omit initial value and amount of increase
	for; sum < 3000; {
		sum += sum
		fmt.Println("second loop: " + strconv.Itoa(sum))
	}

	// can omit ;(semi-colon)
	for sum < 10000 {
		sum += sum
		fmt.Println("third loop: " + strconv.Itoa(sum))
	}

	// can describe forever loop
	for {
		sum += sum
		if sum > 40000 {
			fmt.Println("after fourth loop: " + strconv.Itoa(sum))
			return
		}
	}
}