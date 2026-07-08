package main

import "fmt"

func getOsName() string {
	return "mac"
}

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

	// 関数から値を取得
	os2 := getOsName()
	switch os2 {
	case "mac":
		fmt.Println("Mac!")
	case "Windows":
		fmt.Println("Windows!")
	default: // 条件で一致しない場合
		fmt.Println("default")
	}

	// 特殊な記載の仕方
	switch os3 := getOsName(); os3 {
	case "mac":
		fmt.Println("Mac!")
	case "Windows":
		fmt.Println("Windows!")
	default: // 条件で一致しない場合
		fmt.Println("default")
	}
}
