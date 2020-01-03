package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// condition: F0 = 0, F1 = 1, F(n+2) = Fn + F(n+1) (n ≧ 0)
func fibonacci() func() int {
	val1, val2:= 0, 1
	return func() int{
		val1, val2 = val1+val2, val1
		return val2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
