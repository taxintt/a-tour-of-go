package main

import (
	"fmt"
)

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main(){
	s := []int{7, 2, 8, -9, 4, 0}

	// channelの生成, buffer=0でなぜ普通に利用できている?
	c := make(chan int)

	// 2つのgoroutineの作成+共通したchannelを両方のgoroutineに渡す
	// 送信/受信のどちらかが準備できるまで送受信はブロックされる
	go sum(s[:len(s)/2], c) //[7,2,8] -> 17
	go sum(s[len(s)/2:], c) //[-9,4,0] -> -5

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}