// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

// makeに関しての講義
func main() {
	n := make([]int, 3, 5) //make でスライスの作成
	fmt.Printf("len=%d cap=%d val=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0) // スライスに追加する
	fmt.Printf("len=%d cap=%d val=%v\n", len(n), cap(n), n)
	b := make([]int, 0) // 空のスライスを作成する
	var c []int         // 空のスライスを作成する (メモリを確保しない)
	fmt.Printf("len=%d cap=%d val=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d val=%v\n", len(c), cap(c), c)

	// 演習
	fmt.Println("make でスライスのサイズを明示")
	c = make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
	fmt.Println("make でスライスのサイズを明示しない")

	var d int[]
    d = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}
	fmt.Println(d)
}
