// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)

	var y float32 = 1.2
	yy := int(y)

	fmt.Printf("%T , %v , %f\n", xx, xx, xx)
	fmt.Printf("%T , %v , %d\n", yy, yy, yy)

	// 文字列型の数値を int 型に変換したい(strconv を使用することによって変換可能)
	var s string = "14"
	i, _ := strconv.Atoi(s)

	fmt.Printf("%T , %v , %d\n", i, i, i)
}
