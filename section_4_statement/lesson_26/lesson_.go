package main

import "fmt"

// range に関しての講義
func main() {
	line := []string{"python", "Go", "Java"}

	// 通常のforの記載の仕方
	for i := 0; i < len(line); i++ {
		fmt.Println(i, ":", line[i])
	}
	// rangeの記載の仕方
	for i, v := range line {
		fmt.Println(i, ":", v)
	}
	// インデックス番号は不要の場合の記載の仕方
	for _, v := range line {
		fmt.Println(v)
	}

	// mapに関しても同様に記載が可能
	m := map[string]int{"Apple": 100, "banana": 200, "strawberry": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// map でも同様に片方の値は不要の場合の記載の仕方(値を取り出す)
	for _, v := range m {
		fmt.Println(v)
	}
	// map でも同様に片方の値は不要の場合の記載の仕方(キーを取り出す)
	for k := range m {
		fmt.Println(k)
	}

}
