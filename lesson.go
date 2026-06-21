// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

// 配列に関しての講義
func main() {
	// 配列の初期化
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列に代入する
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
}
