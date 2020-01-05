package main

import (
	"io"
	// "os"
	"strings"
	"fmt"
)

type rot13Reader struct{
	r io.Reader
}

// Read メソッドを実装することで io.Reader インタフェースを満たす
func (r13 rot13Reader) Read(b []byte) (int, error){
	n ,err := r13.r.Read(b)
	for i := range b{
		// substitution cipher
		
	}
	return len(b), nil
}

func main(){
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	// io.Copy(os.stdout, &r)
}