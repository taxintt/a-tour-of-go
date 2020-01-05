package main

import (
	"golang.org/x/tour/reader"
	"strings"
)

type MyReader struct{}

// ASCII文字 'A' の無限ストリームを出力するReader型の実装
// -> 任意のbyte数に対して、全てAで埋めた状態でbyte_size(int)とerrorを返す
func (r *MyReader) Read(b []byte) (int,error){
	for i := range b {
        b[i] = 'A'
	}
	
	// return int, error
    return len(b), nil
}

func main() {
	// readメソッドを実行しなくてもOK(メソッドの実装のみが条件)
	reader.Validate(MyReader{})
}