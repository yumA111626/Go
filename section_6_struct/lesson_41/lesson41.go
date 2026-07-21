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

type Vertex3D struct {
	Vertex // 継承のようなもの
	z      int
}

// func area3D(v Vertex3D) int {
// 	return v.x * v.y * v.z
// }

// ポインタレシーバ
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

// 他言語で言うところのinit処理
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

// // 擬似メソッドの書き方
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// embedについての講義
func main() {
	// v := Vertex{3, 4} // 通常の記載の仕方
	v := New(3, 4, 5)
	// fmt.Println(v.Area())
	v.Scale(10) // 値の書き換え
	fmt.Println(v.Area3D())
}
