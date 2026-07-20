package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func area(v Vertex) int {
	return v.X * v.Y
}

// 擬似メソッドの書き方
func (v Vertex) Area() int {
	return v.X * v.Y
}

// メソッドについての講義
func main() {
	v := Vertex{3, 4}
	fmt.Println(area(v))
	fmt.Println(v.Area())
}
