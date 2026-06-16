// This is Go Main Func
package main // これは必ず記載する

import "fmt" // 論理値型 に関しての講義
func main() {
	// var t ,f bool = true , false
	t, f := true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)

	// 論理演算子
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
}
