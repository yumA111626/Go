package main

import (
	"fmt"
	"os"
)

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

// defer の使用例
func read_file() {
	file, _ := os.Open("/Users/yuma/Go_learn/section_4_statement/lesson_28/lesson_.go") // ファイルを開く
	defer file.Close()                                                                  // 関数実行後にファイルを閉じ忘れを防止できる

	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

func main() {
	// foo()

	// defer fmt.Println("world")
	// fmt.Println("Hello")
	// stacking_defer()
	read_file()
}
