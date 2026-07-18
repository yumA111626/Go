package main

import "fmt"

// 構造体について記載

// 構造体について定義(memberは小文字だと外部からアクセス不可 )
type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v1 := Vertex{X: 1, Y: 2, S: "test"} // 構造体メンバを指定して代入
	fmt.Println(v1)

	v2 := Vertex{} // メンバに値を渡さない -> メンバの方のデフォルト値が入る
	fmt.Println(v2)

	v3 := Vertex{1, 4, "Test"} // member指定して値を渡さない
	fmt.Println(v3)
}
