package main

import (
	"fmt"
	"strconv"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func conv_split(sum int) (y string) {
	//strconv.Itoa: convert int to string、strconv.Atoi: convert string to int
	x := strconv.Itoa(sum * 4 / 9)

	// _: err(type), we don't use error variable
	z, _ := strconv.Atoi(x)
	y = strconv.Itoa(z - 1)
	return 
}

func main() {
	fmt.Println(split(17))
	fmt.Println(conv_split(17))

	x := 1 
	fmt.Println(x)
}
