package main

import (
	"fmt"
	"strings"
)

func main(){
	// Create a tic-tac-toe board(三目並べ)
	board := [][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}

	// the player takes turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// len(board): 3
	for i:=0; i<len(board); i++{
		// join all strings and display row
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
	
}