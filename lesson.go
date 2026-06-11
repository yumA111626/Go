// This is Go Main Func
package main // これは必ず記載する

import "fmt"

// 変数宣言 に関しての講義
func main() {
	var i int = 20
	var t bool = true
	var f bool = false
	var s string = "Test"
	var fl float64 = 1.01

	// 動的型付け宣言
	xi := 1
	xs := "test"
	xf64 := 1.01
	xt, xf := true, false

	fmt.Println("標準的な記載の仕方", i, t, f, s, fl)
	fmt.Println(xi, xs, xf64, xt, xf)
}
