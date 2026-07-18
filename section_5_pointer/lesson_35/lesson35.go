package main

import (
	"fmt"
)

func nil_pointer() {
	// 通常
	var n *int
	fmt.Println(n)
	// ここでインクリメントする -> nil pointer に対して値を追加できない
	*n++
	fmt.Println(n)
}

// make とnew の違いについて
func diff_make_new() {
	// make
	s := make([]int, 0)
	fmt.Println(s)

	m := make(map[string]int)
	fmt.Println(m)

	// 構造体
	var st = new(struct{})
	fmt.Printf("%T\n", st)

	// チャネル
	var ch = new(chan int)
	fmt.Printf("%T\n", ch)
}

// new , make の違いについて
func main() {
	// 値を格納しないがメモリの領域を確保したい

	// 通常
	var n *int
	fmt.Println(n) //  <nil>

	// メモリを確保する記載の仕方(new を記載する)
	var n_2 *int = new(int)
	fmt.Println(n_2)

	// nil_pointer()
	diff_make_new()
}
