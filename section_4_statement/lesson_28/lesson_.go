package main

import "fmt"

/*
	defer に関して記載

defer とは遅延実行
*/

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("Hello foo")
}

func main() {
	foo()

	defer fmt.Println("world")
	fmt.Println("Hello")
}
