package main

import (
	"io"
	"log"
	"os"
)

// logging に関して記載
/*
他の言語でいうところの info , warm , exception などは標準のものでは存在しない
*/

// logging 設定をする
func loggingSettings(write_log_file string) {
	log_file, _ := os.OpenFile(write_log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, log_file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

// 存在しないファイルを開こうとしてエラーを発生させる
func read_file() {
	_, err := os.Open("sasasas")
	if err != nil {
		log.Fatalln("Exit", err)
	}
}
func main() {
	// log.Println("Info!")
	// log.Printf("%T , %v", "Test", "Test")
	// log.Fatalln("Error!!") // fatal は実行後にプログラムが終了してしまう
	loggingSettings("test.log")
	read_file()
}
