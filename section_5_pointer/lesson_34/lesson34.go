package main

import "fmt"

// ポインタについての講義
func main() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
}
