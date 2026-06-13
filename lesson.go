// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt" // TODO: git commit コメントにセクションNoとタイトルを記載すること
	"strings"
)

// 文字列型 に関しての講義
func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + " World")       // 分割して表示
	fmt.Println("Hello World"[0])         // Python ならスライスで指定して特定の文字を取得可能 (Go の場合 ASCII コード :72)
	fmt.Println(string("Hello World"[0])) // Stringにキャストすればpython のように取得可能 : H

	// リテラルの変更 (特定の文字のみ変更したいなど)
	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1)) // : Xello World 仕組みとしては、sをコピーして変更するため元の値は変更されない
	fmt.Println(s)

	// 特定の文字を含んでいるかの判定
	fmt.Println(strings.Contains(s, "World"))

	// 文字列結合
	fmt.Println("Hello" +
		"World")

	// " を表示させたい
	fmt.Println("\"")
	fmt.Println(`"`)

}
