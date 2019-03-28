package basics_1

import (
	"fmt"
	"strconv"
)

// func split(sum int) (x, y string) {
// 	x = strconv.Itoa(sum * 4 / 9)
// 	y = strconv.Itoa(sum - 1)
// 	return
// }
func split(sum int) (y string) {
	//関数内では := の代入文を利用可能
	x := strconv.Itoa(sum * 4 / 9)
	z, _ := strconv.Atoi(x) //return two values
	y = strconv.Itoa(z - 1)
	return y 
}

func main() {
	//math.pi と math.Pi の違い = 先頭が大文字はモジュール外部から参照可能
	fmt.Println(split(17))

	//var x int = 1 
	x := 1 
	fmt.Println(x)
}
