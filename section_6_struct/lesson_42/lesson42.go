package main

import "fmt"

// typedef みたいなもの
type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

// non-struct に関しての講義
func main() {
	myInt := MyInt(4)
	fmt.Println(myInt.Double())
}
