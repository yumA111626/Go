package main

import "fmt"

// for 文に関しての講義
func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("Continue")
			continue // 次のfor にいく
		}

		if i > 5 {
			fmt.Println("Break")
			break
		}
	}
	// 他のfor 文の記載の仕方
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	// 作為的に無限ループ生成する
	for {
		fmt.Println("Hello")
	}
}
