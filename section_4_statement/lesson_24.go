package main

import "fmt"

/*
==============
if 文に関しての講義
==============
*/

func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "NG"
	}
}

func main() {
	x := 9
	if x%2 == 0 {
		fmt.Println("By 2")
	} else if x%3 == 0 {
		fmt.Println("By 3")
	} else {
		fmt.Println("else")
	}

	//偶数かを判定する
	jug := by2(3)
	if jug == "OK" {
		fmt.Println("Great")
	} else {
		fmt.Println("Oh...")
	}

	// if ステートメントの中でしか使用しない変数の場合の書き方
	if result := by2(10); result == "OK" {
		fmt.Println("Great OK")
	} else {
		fmt.Println("Oh.......")
	}
}
