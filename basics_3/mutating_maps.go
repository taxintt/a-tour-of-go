package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// can confirm whether the key exists or not
	// do not forget :=
	val, ok := m["Answer"]
	fmt.Println("The value:", val, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// can confirm whether the key exists or not
	val, ok = m["Answer"]
	fmt.Println("The value:", val, "Present?", ok)
}