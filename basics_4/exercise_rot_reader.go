package main

import (
	"io"
	"os"
	"strings"
)

// wrap io.Reader
type rot13Reader struct{
	r io.Reader
}

// Read メソッドを実装することで io.Reader インタフェースを満たす
func (r13 rot13Reader) Read(b []byte) (int, error){
	//
	n ,err := r13.r.Read(b)

	for i := range b{
		// substitution cipher
		switch {
		case b[i] >= 'A' && b[i] <= 'M' || b[i] >= 'a' && b[i] <= 'm':
			b[i] += 13
		case b[i] >= 'N' && b[i] <= 'Z' || b[i] >= 'n' && b[i] <= 'z':
			b[i] -= 13
		}
	}
	return n ,err
}

func main(){
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}