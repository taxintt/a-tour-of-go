package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	walk(t, ch)
	close(ch)
}

// サブ的に関数を作成、再帰的に関数を呼び出してchannelにvalueを放り込む
func walk(t *tree.Tree, ch chan int){
	if t.Left != nil{
		walk(t.Left, ch)
	}
	if t.Right != nil{
		walk(t.Right, ch)
	}
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for{
		v1, ok1 := <-c1
        v2, ok2 := <-c2

		switch {
		// t1 and t2 contain the same values -> 一部が同じ値であればOK
		// multiple casesではコンマ区切りで記載
        case !ok1, !ok2:
            return ok1 == ok2
        case v1 != v2:
            return false
        }
	}
	
}

func main() {
	k := 1 
	ch := make(chan int, 10)

	go Walk(tree.New(k), ch)
	for i := range ch {
        fmt.Println(i)
    }

	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}