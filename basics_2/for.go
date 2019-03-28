package basics_2

import "fmt"

func main(){
	sum := 0
	for i:=0;  i<10;  i++{
		sum += 1
	}
	// セミコロンの省略は下の形式でのみ可能
	// for  sum < 1000 {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }
	fmt.Println(sum)
}