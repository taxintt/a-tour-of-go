package basics_1

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
const(
	//2進数の１を100桁ずらす
	normal = 1 << 2
	Big = 1 << 100
	Small = Big >> 99
)


func needInt(x int) int {
	return x*10 + 1 
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	//%T：対象データの型情報を埋め込む
	//%v：デフォルトフォーマットで対象データの情報を埋め込む
	//%+v：構造体を出力する際に、%vの内容にフィールド名も加わる
	//%#v：Go言語のリテラル表現で対象データの情報を埋め込む
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println(normal)
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
