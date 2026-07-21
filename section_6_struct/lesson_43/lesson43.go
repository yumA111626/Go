package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

// インターフェースに関しての講義
func main() {
	var Mike Human = Person{"Mike"}
	Mike.Say()
}
