package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func area(v Vertex) int {
	return v.X * v.Y
}

// ポインタレシーバ
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// 擬似メソッドの書き方
func (v Vertex) Area() int {
	return v.X * v.Y
}

// メソッドについての講義
func main() {
	v := Vertex{3, 4}
	// fmt.Println(area(v))
	fmt.Println(v.Area())
	v.Scale(10) // 値の書き換え
	fmt.Println(v.Area())
}
