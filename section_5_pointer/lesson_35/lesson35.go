package main

import "fmt"

// new , make の違いについて
func main() {
	// 値を格納しないがメモリの領域を確保したい

	// 通常
	var n *int
	fmt.Println(n) //  <nil>

	// メモリを確保する記載の仕方(new を記載する)
	var n_2 *int = new(int)
	fmt.Println(n_2)
}
