// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

func add(x int, y int) {
	fmt.Println(x + y)
}

// 関数に関しての講義
func main() {
	add(10, 20)
}
