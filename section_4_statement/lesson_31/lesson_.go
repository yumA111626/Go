package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to access DB!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

// panic , recover について記載
func main() {
	save()
}
