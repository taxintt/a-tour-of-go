package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct{
	v map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter)Inc(key string){
	c.mux.lock() //1

	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++ //2
	c.mux.unlock() //3
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter)Value(key string)int{
	c.mux.lock() //1
	
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.unlock() //3
	return c.v[key] //2
}

func main(){
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i<1000; i++{
		go c.Inc("somekey")
	}

	time.Sleep(time.second)
	fmt.Println(c.Value("somekey"))
}