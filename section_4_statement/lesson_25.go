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
		fmt.Println(i)
	}
}
