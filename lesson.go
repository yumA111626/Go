// This is Go Main Func
package main // これは必ず記載する

import "fmt"

const Pi = 3.14

const (
	USER_NAME = "user_name"
	PASS_WORD = "Pass_word"
)

const Int_Max = 9223372036854775807 + 1

// var max int = 9223372036854775807 + 1 // 変数の場合、コンパイラに Intと判断され、コンパイルエラーになる

// 変数宣言 に関しての講義
func main() {
	fmt.Println(Pi, USER_NAME, PASS_WORD)
	fmt.Println(Int_Max - 1) // const の場合はコンパイル時には型の判断がされない
}
