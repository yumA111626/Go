package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func area(v Vertex) int {
	return v.x * v.y
}

// ポインタレシーバ
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// // 擬似メソッドの書き方
func (v Vertex) Area() int {
	return v.x * v.y
}

// 他言語で言うところのinit処理
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

// コンストラクタについての講義
func main() {
	// v := Vertex{3, 4} // 通常の記載の仕方
	v := New(3, 4)
	// fmt.Println(v.Area())
	v.Scale(10) // 値の書き換え
	fmt.Println(v.Area())
}
