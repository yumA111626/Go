// This is Go Main Func
package main // これは必ず記載する

import "fmt" // 型変換 に関しての講義

func incrementGenerator() func() int {
	x := 0
	// インナー関数を記載
	return func() int {
		x++
		return x
	}
}

// クロージャーに関しての講義
func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
