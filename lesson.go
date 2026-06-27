// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // 型変換 に関しての講義
)

// mapに関しての講義
func main() {
	// map とはpython でいう所の辞書のようなもの
	m := map[string]int{"apple": 200, "banana": 100}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["melon"] = 250 // map に追加をする
	fmt.Println(m)
	fmt.Println(m["melon"])

	// 存在しないものを出力する
	fmt.Println(m["nothing"])

	// 存在しているかの判定
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// メモリ上にからのmapを作成してそこのmapに追加していく
	m2 := make(map[string]int) // map を作成
	m2["PC"] = 3000
	fmt.Println(m2)

	// var での宣言時の注意点
	var m3 map[string]int
	m3["PC"] = 5000
	fmt.Println(m3)
}
