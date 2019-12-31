package main

import (
	"fmt"
	"time"
	"reflect"
)

func main(){
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	fmt.Println(reflect.TypeOf(today))  

	// conditions(switch <variable>)
	switch time.Saturday {

	// today + i -> invalid operationとなるがなぜ?
	case today+0:
		fmt.Println("Today")
	case today+1:
		fmt.Println("Tommorow")
	case today+2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}