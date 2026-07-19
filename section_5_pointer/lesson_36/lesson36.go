package main

import "fmt"

// 構造体について記載

// 構造体について定義(memberは小文字だと外部からアクセス不可 )
type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)
	fmt.Println(*v2)

	/*
		v1 := Vertex{X: 1, Y: 2, S: "test"} // 構造体メンバを指定して代入
		fmt.Println(v1)

		v2 := Vertex{} // メンバに値を渡さない -> メンバの方のデフォルト値が入る
		fmt.Println(v2)

		v3 := Vertex{1, 4, "Test"} // member指定して値を渡さない
		fmt.Println(v3)

		var v4 Vertex // typedef みたいな記載の仕方
		fmt.Println(v4)
		v4.X = 1
		fmt.Println(v4)

		// 構造体pointer
		v5 := new(Vertex)
		fmt.Println(v5)

		v6 := &Vertex{}
		fmt.Println(v6)
	*/
}
