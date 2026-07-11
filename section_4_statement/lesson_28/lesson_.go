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

// stacking_defer に関して記載
// したから順に実行されていく
func stacking_defer() {
	fmt.Println("Run")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("Success")

}

func main() {
	// foo()

	// defer fmt.Println("world")
	// fmt.Println("Hello")
	stacking_defer()
}
