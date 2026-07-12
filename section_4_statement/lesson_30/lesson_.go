package main

import (
	"fmt"
	"log"
	"os"
)

// エラーハンドリングに関する講義
func main() {
	file, err := os.Open("./section_4_statement/lesson_30/lesson_.go")
	if err != nil {
		log.Fatalln("Error")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}

	fmt.Println(count, string(data))
}
