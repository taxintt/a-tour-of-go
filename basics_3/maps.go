package main

import (
	"fmt"
)

// obj's name is Vertex
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main(){
	m = make(map[string]Vertex)

	// map key and value
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	// display value using key
	fmt.Println(m["Bell Labs"])
}