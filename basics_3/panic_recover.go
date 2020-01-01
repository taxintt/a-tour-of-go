package main

import (
	"fmt"
)

func recover_panic(){
	defer func(){
		fmt.Println("End")
		// recover: stop current function ,and return error message (スタックメモリの巻き戻し(解放?)は行わない)
		err := recover()
		if err != nil {
			fmt.Println("Recover!: ", err)
		}
		fmt.Println("recover: end")

	}()
	fmt.Println("Start")

	// cause panic(=runtime problem) - ex) index out of range
	text := "hoge"
	fmt.Println("2:", text[5:])
	
	// caution: didn't print this message
	fmt.Println("not printed")
}

func main(){
	recover_panic()
	fmt.Println("recover_panic: end")
}