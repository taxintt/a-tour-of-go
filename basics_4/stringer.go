package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

// method - value receiver
// String() - method that is inherited by interface(stringer)
func (p Person) String() string{
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main(){
	a := Person{Name: "Nicholas Cage", Age: 21}
	b := Person{Name: "Zaphod Beeblebrox", Age: 9001}

	// fmt package has String() interface
	fmt.Println(a,b)
}