package main

import "fmt"

// switch に関して記載
func main() {
	// switch 文の記載の仕方
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "Windows":
		fmt.Println("Windows!")
	default: // 条件で一致しない場合
		fmt.Println("default")
	}
}
