package main

import "fmt"

// range に関しての講義
func main() {
	line := []string{"python", "Go", "Java"}

	for i := 0; i < len(line); i++ {
		fmt.Println(i, ":", line[i])
	}
}
