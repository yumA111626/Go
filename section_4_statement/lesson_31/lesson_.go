package main

func thirdPartyConnectDB() {
	panic("Unable to access DB!")
}

func save() {
	thirdPartyConnectDB()
}

// panic , recover について記載
func main() {
	save()
}
