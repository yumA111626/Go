package main

import "log"

// logging に関して記載
/*
他の言語でいうところの info , warm , exception などは標準のものでは存在しない
*/
func main() {
	log.Println("Info!")
	log.Printf("%T , %v", "Test", "Test")
	log.Fatalln("Error!!") // fatal は実行後にプログラムが終了してしまう
}
