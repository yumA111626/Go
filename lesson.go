// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

// スライスに関しての講義
func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:2])

	// 配列の中身の変更
	n[2] = 100
	fmt.Println(n)
	fmt.Println(n[2])

	// ２次元配列
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	// 配列に追加
	n = append(n, 300)
	fmt.Println(n)
}
