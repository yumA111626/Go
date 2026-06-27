// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

// byteに関しての講義
func main() {
	by := []byte{72, 73} // ASCII Code
	fmt.Println(by)

	by_str := []byte("HI")
	fmt.Println(by_str)
	fmt.Println(string(by_str))
}
