// This is Go Main Func
package main // これは必ず記載する

import "fmt"

// 短変数宣言は関数内でしか実行できない var の通常宣言は可能
var (
	i  int     = 20
	t  bool    = true
	f  bool    = false
	s  string  = "Test"
	fl float64 = 1.01
)

func foo() {
	xi := 1
	xs := "test"
	xf64 := 1.01
	xt, xf := true, false

	fmt.Println(xi, xs, xf64, xt, xf)
	fmt.Printf("%T\n", xf64) // 型について出力させる
}

// 変数宣言 に関しての講義
func main() {
	foo() // 関数内では短変数宣言は可能
	// fmt.Println("標準的な記載の仕方", i, t, f, s, fl)
}
