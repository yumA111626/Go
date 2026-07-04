package main

import "fmt"

/*
==============
if 文に関しての講義
==============
*/
func main() {
	x := 9
	if x%2 == 0 {
		fmt.Println("By 2")
	} else if x%3 == 0 {
		fmt.Println("By 3")
	} else {
		fmt.Println("else")
	}
}
