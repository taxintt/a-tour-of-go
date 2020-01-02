package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	fmt.Println(pic)

	for y := range(pic) {

		// my answer 
		tmp := make([]uint8, dx)
		for x := range(tmp) {
			tmp[x] = uint8(x*y)
		}
		pic[y] = tmp

		// pic[y] = make([]uint8, dx)
		// for x := range(pic[y]){
		// 	pic[y][x] = uint8(x*y)
		// }
	}

	return pic
}

func main() {
	pic.Show(Pic)
}