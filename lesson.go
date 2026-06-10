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

	// まとめて宣言
	var (
		num       int     = 10
		is_ok     bool    = true
		is_not_ok bool    = false
		str_test  string  = "This is Test"
		fl_num    float64 = 1.02
	)

	fmt.Println("標準的な記載の仕方", i, t, f, s, fl)
	fmt.Println("まとめて宣言をした記載の仕方", num, is_ok, is_not_ok, str_test, fl_num)
}
