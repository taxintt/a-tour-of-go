package main 

import (
	"fmt"
	"image"
)

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func main(){
	// create RGBA from Rect(イメージの範囲)
	// x: 0, y: 0, width: 100, height:100
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Bounds: 画像の領域を取得する
	fmt.Println(m.Bounds())

	// (0,0)のRGBA(color)を返す
	fmt.Println(m.At(0, 0).RGBA())
}