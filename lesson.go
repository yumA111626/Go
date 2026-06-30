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

// 円の面積を算出する関数(クロージャー)を記載
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// クロージャーに関しての講義
func main() {
	c1 := circleArea(3.14)
	fmt.Println(c1(3))
	c2 := circleArea(3)
	fmt.Println(c2(3))
}
