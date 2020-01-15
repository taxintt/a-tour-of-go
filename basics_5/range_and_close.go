package main

import (
	"fmt"
)

func fibonacci(n int, c chan int){
	x, y := 0, 1
	for i:=0; i<n; i++{
		// send to channel
		c <- x
		x,y = y, x+y
	}
	// close the channel
	// after close, we can't send data to channel
	close(c)

	// happen panic because the channel is closed
	// c <- x
}

func main(){
	c := make(chan int, 10)

	// channelにデータを詰める
	go fibonacci(cap(c), c)

	// closeによって送信を止めてから、データを取り出していく
	// closeしないと待ち状態が継続してしまうのでdeadlockのエラーが発生する?
	for i := range c{
		fmt.Println(i)
	}

	// for文でchannelから全ての値を受信したのでchannelのopen(bool)がfalseになる
	_, ok := <-c
	fmt.Println(ok)
}