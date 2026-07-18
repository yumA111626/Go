package main

import "fmt"

// 普通の関数
func one(x int) {
	x = 1
}

// 普通の関数
func one_pointer(x *int) {
	*x = 1
}

// ポインタについての講義
func main() {
	var n int = 100
	one(n)
	fmt.Println(n)
	one_pointer(&n)
	fmt.Println(n)
	/*
		fmt.Println(n)
		fmt.Println(&n)

		var p *int = &n
		fmt.Println(p)
		fmt.Println(*p)
	*/
}
