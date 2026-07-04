// This is Go Main Func
package main // これは必ず記載する

import "fmt" // 型変換 に関しての講義

/*
可変長の引数を渡せる関数を作成する場合は、...をつける
*/
func foo(params ...int) {
	fmt.Println(len(params), params)
	// ループで引数の中身を取り出す
	for _, param := range params {
		fmt.Println(param)
	}
}

// 可変長引数に関しての講義
func main() {
	foo(10, 20)
	foo(10, 20, 30)
}
