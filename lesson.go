// This is Go Main Func
package main // これは必ず記載する

import "fmt" // 型変換 に関しての講義

/*
	同じ型の場合は、引数に関しても省略して記載することは可能

func add(x,y int) int{}
*/
func add(x int, y int) int {
	return x + y
}

// 返り値が複数存在する
func add_2(x, y int) (int, int) {
	return x + y, x - y
}

// 返り値で変数宣言をする
func cal(price, item int) (result int) {
	result = price * item
	return result // 変数で返り値を宣言しているため return のみでも良い
}

// クロージャーに関しての講義
func main() {
	x := 0
	// インナー関数を記載
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
