package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// don't forget initialization
	m := make(map[string]int)
	arr := strings.Fields(s)
	
	for i, _ := range(arr) { 
		// increment value in a map - store["key"] += num
		m[arr[i]] += 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
